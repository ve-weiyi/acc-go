package bean

import "acc/app/model/entity"

type LoginResp struct {
	Uid      int    `json:"uid" example:"1"`
	Username string `json:"username" example:"admin"`
	Details  *entity.UserInfo
	Token    string `json:"token" example:""`
}
