package routes

import (
	"eco_points/config"
	"eco_points/internal/features/locations"
	"eco_points/internal/features/trashes"
	users "eco_points/internal/features/users"
	deposits "eco_points/internal/features/waste_deposits"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.Handler, th trashes.HandlerTrashInterface, dh deposits.HandlerInterface, l locations.HandlerInterface) {

	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())

	UsersRoute(e, uh)
	TrashRoute(e, th, dh)
	LocRoute(e, l)
}

func UsersRoute(e *echo.Echo, uh users.Handler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.GET("", uh.GetUser())
	u.PUT("", uh.UpdateUser())
	u.DELETE("", uh.DeleteUser())
}

func TrashRoute(e *echo.Echo, th trashes.HandlerTrashInterface, dh deposits.HandlerInterface) {
	t := e.Group("/trash")
	t.Use(JWTConfig())
	t.POST("", th.AddTrash())
	t.GET("", th.GetTrash())
	t.DELETE("/:id", th.DeleteTrash())

	d := e.Group("/deposit")
	d.Use(JWTConfig())
	d.POST("", dh.DepositTrash())
	d.PUT("", dh.UpdateWasteDepositStatus())
	d.GET("", dh.GetUserDeposit())
}

func LocRoute(e *echo.Echo, l locations.HandlerInterface) {
	t := e.Group("/location")
	t.Use(JWTConfig())
	t.POST("", l.AddLocation())

}
func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
