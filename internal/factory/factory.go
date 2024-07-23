package factory

import (
	"eco_points/config"

	u_hnd "eco_points/internal/features/users/handler"
	u_rep "eco_points/internal/features/users/repository"
	u_srv "eco_points/internal/features/users/service"
	"eco_points/internal/routes"
	"eco_points/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	db, _ := config.ConnectDB()

	pu := utils.NewPassUtil()
	tu := utils.NewJwtUtility()

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserService(uq, pu, tu)
	uh := u_hnd.NewUserHandler(us, tu)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh)
}
