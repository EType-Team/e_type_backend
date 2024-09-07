package router

import (
	"api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, uwpc controller.IUserWordProgressController) *echo.Echo {
	e := echo.New()
	a := e.Group("/api")
	a.POST("/signup", uc.SignUp)
	a.POST("/login", uc.LogIn)
	a.POST("/logout", uc.LogOut)
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
