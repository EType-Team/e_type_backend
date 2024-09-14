package router

import (
	"api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
	lc controller.ILessonController,
	lwc controller.ILessonWordController,
	uwpc controller.IUserWordProgressController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken,
			"Authorization", "Cookies"},
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

	securedGroup := a.Group("")
	securedGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	l := securedGroup.Group("/lessons")
	l.GET("", lc.GetAllLesson)
	l.GET("/:lessonId", lc.GetLessonById)

	lw := securedGroup.Group("/lessonWord")
	lw.GET("/:lessonId", lwc.GetLessonWordByLessonId)

	uwp := securedGroup.Group("/userWordProgresses")
	uwp.GET("", uwpc.GetAllUserWordProgress)
	uwp.GET("/:userWordProgressId", uwpc.GetUserWordProgressById)
	uwp.POST("", uwpc.CreateUserWordProgress)
	uwp.PUT("/:userWordProgressId", uwpc.UpdateUserWordProgress)
	return e
}
