package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func IsAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "認証エラーです",
			})
		}
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "管理者権限が必要です",
			})
		}
		return next(c)
	}
}
