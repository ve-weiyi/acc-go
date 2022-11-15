package dto

type AuthReq struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123456"`
	Code     string `json:"code" example:""`
}
