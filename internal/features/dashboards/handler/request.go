package handler

type StatusRequest struct {
	Status string `json:"status" form:"status"`
}
