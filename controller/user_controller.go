package controller

import (
	"api/model"
	"api/usecase"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type IUserController interface {
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	CsrfToken(c echo.Context) error
	GoogleLogin(c echo.Context) error
	GoogleCallback(c echo.Context) error
	Logout(c echo.Context) error
}

type userController struct {
	uu          usecase.IUserUsecase
	oauthConfig *oauth2.Config
	adminEmails []string
}

func NewUserController(uu usecase.IUserUsecase, oauthConfig *oauth2.Config) IUserController {
	adminEmailsStr := os.Getenv("ADMIN_EMAILS")
	adminEmails := strings.Split(adminEmailsStr, ",")
	return &userController{
		uu,
		oauthConfig,
		adminEmails,
	}
}

func (uc *userController) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userRes, err := uc.uu.GetUserById(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userRes, err := uc.uu.UpdateUser(user, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc *userController) GoogleLogin(c echo.Context) error {
	url := uc.oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (uc *userController) GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "state-token" {
		return c.JSON(http.StatusBadRequest, "Invalid state parameter")
	}

	code := c.QueryParam("code")
	token, err := uc.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Code exchange failed")
	}

	client := uc.oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get user info")
	}
	defer userInfoResp.Body.Close()

	var userInfo struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse user info")
	}

	user := model.User{
		Email: userInfo.Email,
		Name:  userInfo.Name,
		Image: userInfo.Picture,
		Role:  "user",
	}

	for _, adminEmail := range uc.adminEmails {
		if user.Email == strings.TrimSpace(adminEmail) {
			user.Role = "admin"
			break
		}
	}
	storedUser := model.User{}
	err = uc.uu.GetUserByEmail(&storedUser, user.Email)
	if err != nil {
		err = uc.uu.CreateUser(&user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create user")
		}
		storedUser = user
	}

	jwtToken, err := uc.uu.GenerateJWT(storedUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwtToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_REDIRECT_URL"))
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
