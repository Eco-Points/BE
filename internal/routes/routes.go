package routes

import (
	"eco_points/config"
	users "eco_points/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.Handler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())

	UsersRoute(e, uh)
}

func UsersRoute(e *echo.Echo, uh users.Handler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.GET("", uh.GetUser())
	u.PUT("", uh.UpdateUser())
	u.DELETE("", uh.DeleteUser())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
