package handler

type UserDashboardResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsAdmin  bool   `json:"is_admin"`
	Status   string `json:"status"`
	Point    uint   `json:"point"`
}
