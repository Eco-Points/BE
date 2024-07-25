package factory

import (
	"eco_points/config"

	u_hnd "eco_points/internal/features/users/handler"
	u_rep "eco_points/internal/features/users/repository"
	u_srv "eco_points/internal/features/users/service"

	t_hnd "eco_points/internal/features/trashes/handler"
	t_rep "eco_points/internal/features/trashes/repository"
	t_srv "eco_points/internal/features/trashes/service"

	d_hnd "eco_points/internal/features/waste_deposits/handler"
	d_rep "eco_points/internal/features/waste_deposits/repository"
	d_srv "eco_points/internal/features/waste_deposits/service"

	"eco_points/internal/routes"
	"eco_points/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	db, _ := config.ConnectDB()

	pu := utils.NewPassUtil()
	jwt := utils.NewJwtUtility()

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserService(uq, pu, jwt)
	uh := u_hnd.NewUserHandler(us, jwt)

	tq := t_rep.NewTrashQuery(db)
	tu := t_srv.NewTrashService(tq)
	th := t_hnd.NewTrashHandler(tu, jwt)

	dq := d_rep.NewTrashQuery(db)
	du := d_srv.NewDepositsService(dq)
	dh := d_hnd.NewDepositService(du, jwt)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, th, dh)
}
