package router

import (
	"api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, uwpc controller.IUserWordProgressController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60
	}))
	a := e.Group("/api")
	a.POST("/signup", uc.SignUp)
	a.POST("/login", uc.LogIn)
	a.POST("/logout", uc.LogOut)
	a.GET("/csrf", uc.CsrfToken)
	t := a.Group("/userWordProgresses")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", uwpc.GetAllUserWordProgress)
	t.GET("/:userWordProgressId", uwpc.GetUserWordProgressById)
	t.POST("", uwpc.CreateUserWordProgress)
	t.PUT("/:userWordProgressId", uwpc.UpdateUserWordProgress)
	return e
}
