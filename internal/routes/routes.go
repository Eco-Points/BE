package routes

import (
	"eco_points/config"
	"eco_points/internal/features/dashboards"
	"eco_points/internal/features/excelize"
	"eco_points/internal/features/exchanges"
	"eco_points/internal/features/feedbacks"
	"eco_points/internal/features/locations"
	"eco_points/internal/features/rewards"
	"eco_points/internal/features/trashes"

	users "eco_points/internal/features/users"
	deposits "eco_points/internal/features/waste_deposits"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, th trashes.HandlerTrashInterface, dh deposits.HandlerDepoInterface, l locations.HandlerLocInterface, rh rewards.RHandler, exh exchanges.ExHandler, dsh dashboards.DshHandler, excl excelize.HandlerMakeExcelInterface, fh feedbacks.FHandler) {

	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())

	UsersRoute(e, uh)
	TrashRoute(e, th, dh)
	LocRoute(e, l)
	DashboardRoute(e, dsh)
	RewardRoute(e, rh)
	ExchangeRoute(e, exh)
	excelMakeRoute(e, excl)
	FeedbackRoute(e, fh)
}

func UsersRoute(e *echo.Echo, uh users.UHandler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.GET("", uh.GetUser())
	u.PUT("", uh.UpdateUser())
	u.DELETE("", uh.DeleteUser())
}

func TrashRoute(e *echo.Echo, th trashes.HandlerTrashInterface, dh deposits.HandlerDepoInterface) {
	t := e.Group("/trash")
	t.Use(JWTConfig())
	t.POST("", th.AddTrash())
	t.GET("", th.GetTrash())
	t.DELETE("/:id", th.DeleteTrash())
	t.PUT("/:id", th.UpdateTrash())

	d := e.Group("/deposit")
	d.Use(JWTConfig())
	d.POST("", dh.DepositTrash())
	d.PUT("", dh.UpdateWasteDepositStatus())
	d.GET("", dh.GetUserDeposit())
	d.GET("/:id", dh.GetDepositbyId())

}

func LocRoute(e *echo.Echo, l locations.HandlerLocInterface) {
	t := e.Group("/location")
	t.Use(JWTConfig())
	t.POST("", l.AddLocation())
	t.GET("", l.GetLocation())
}

func DashboardRoute(e *echo.Echo, dsh dashboards.DshHandler) {
	ds := e.Group("/dashboard")
	ds.Use(JWTConfig())
	ds.GET("", dsh.GetDashboard())
	ds.GET("/users", dsh.GetAllUsers())
	ds.GET("/users/:target_id", dsh.GetUser())
	ds.PUT("/users/:target_id", dsh.UpdateUserStatus())
	ds.DELETE("/users/:target_id", dsh.DeleteUserByAdmin())
	ds.GET("/depositstat", dsh.GetDepositStat())
	ds.GET("/rewardstat", dsh.GetRewardStatData())
}

func RewardRoute(e *echo.Echo, rh rewards.RHandler) {
	t := e.Group("/reward")
	t.Use(JWTConfig())
	t.POST("", rh.AddReward())
	t.PUT("/:id", rh.UpdateReward())
	t.DELETE("/:id", rh.DeleteReward())

	e.GET("/reward/:id", rh.GetRewardByID())
	e.GET("/reward", rh.GetAllRewards())
}

func ExchangeRoute(e *echo.Echo, exh exchanges.ExHandler) {
	t := e.Group("/exchange")
	t.Use(JWTConfig())
	t.POST("", exh.AddExchange())
	t.GET("", exh.GetExchangeHistory())
}

func FeedbackRoute(e *echo.Echo, fh feedbacks.FHandler) {
	t := e.Group("/feedback")
	t.Use(JWTConfig())
	t.POST("", fh.AddFeedback())
	t.PUT("/:id", fh.UpdateFeedback())
	t.DELETE("/:id", fh.DeleteFeedback())

	e.GET("/feedback/:id", fh.GetFeedbackByID())
	e.GET("/feedback", fh.GetAllFeedbacks())
}

func excelMakeRoute(e *echo.Echo, excl excelize.HandlerMakeExcelInterface) {
	t := e.Group("/excel")
	t.Use(JWTConfig())
	t.GET("", excl.MakeExcel())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
