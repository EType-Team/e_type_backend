package router

import (
	"api/controller"
	customMiddleware "api/middleware"
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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", http.MethodOptions},
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
	e.GET("/auth/google/login", uc.GoogleLogin)
	e.GET("/auth/google/callback", uc.GoogleCallback)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	l := e.Group("/lessons")
	l.GET("", lc.GetAllLesson)
	l.GET("/:lessonId", lc.GetLessonById)
	l.POST("", lc.CreateLesson)
	l.PUT("/:lessonId", lc.UpdateLesson)
	l.DELETE("/:lessonId", lc.DeleteLesson)
	l.POST("/:lessonId/words", lc.AddNewWordToLesson)
	l.DELETE("/:lessonId/words/:wordId", lc.RemoveWordFromLesson)

	lw := e.Group("/lessonWord")
	lw.GET("/:lessonId", lwc.GetLessonWordByLessonId)

	securedGroup := e.Group("")
	securedGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	adminGroup := securedGroup.Group("/admin")
	adminGroup.Use(customMiddleware.IsAdminMiddleware)

	adminGroup.GET("/lessons", lc.GetAllLesson)
	adminGroup.GET("/lessons/:lessonId", lc.GetLessonById)
	adminGroup.POST("/lessons", lc.CreateLesson)
	adminGroup.PUT("/lessons/:lessonId", lc.UpdateLesson)
	adminGroup.DELETE("/lessons/:lessonId", lc.DeleteLesson)
	adminGroup.POST("/lessons/:lessonId/words", lc.AddNewWordToLesson)
	adminGroup.DELETE("/lessons/:lessonId/words/:wordId", lc.RemoveWordFromLesson)

	user := securedGroup.Group("/user")
	user.GET("", uc.GetUser)
	user.POST("/update", uc.UpdateUser)

	uwp := securedGroup.Group("/userWordProgresses")
	uwp.GET("", uwpc.GetAllUserWordProgress)
	uwp.POST("/incrementProgress", uwpc.IncrementOrCreateUserWordProgress)
	uwp.POST("/incrementTestProgress", uwpc.IncrementOrCreateUserWordTestProgress)
	uwp.GET("/:wordId", uwpc.GetUserWordProgressByWordId)
	uwp.GET("/:lessonId", uwpc.GetUserWordProgressByLessonId)
	return e
}
