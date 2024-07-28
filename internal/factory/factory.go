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

	l_hnd "eco_points/internal/features/locations/handler"
	l_rep "eco_points/internal/features/locations/repository"
	l_srv "eco_points/internal/features/locations/service"

	r_hnd "eco_points/internal/features/rewards/handler"
	r_rep "eco_points/internal/features/rewards/repository"
	r_srv "eco_points/internal/features/rewards/service"

	"eco_points/internal/routes"
	"eco_points/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	db, _ := config.ConnectDB()

	pu := utils.NewPassUtil()
	jwt := utils.NewJwtUtility()
	cloud := utils.NewCloudinaryUtility()

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserService(uq, pu, jwt, cloud)
	uh := u_hnd.NewUserHandler(us, jwt)

	tq := t_rep.NewTrashQuery(db)
	tu := t_srv.NewTrashService(tq, cloud)
	th := t_hnd.NewTrashHandler(tu, jwt)

	dq := d_rep.NewDepoQuery(db)
	dq.SetDbSchema(config.ImportSetting().Schema)

	du := d_srv.NewDepositsService(dq)
	dh := d_hnd.NewDepositHandler(du, jwt)

	lq := l_rep.NewLocQuery(db)
	lu := l_srv.NewLocService(lq)
	lh := l_hnd.NewLocHandler(lu, jwt)

	rq := r_rep.NewRewardModel(db)
	rs := r_srv.NewRewardService(rq)
	rh := r_hnd.NewRewardHandler(rs, jwt, cloud)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, th, dh, lh, rh)
}
