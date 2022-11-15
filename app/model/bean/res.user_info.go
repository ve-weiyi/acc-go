package bean

import "acc/app/model/entity"

type UserInfoResp struct {
	Id       int              `json:"uid"`
	Username string           `json:"username"`
	Roles    []string         `json:"roles"`
	Details  *entity.UserInfo `json:"details"`
}
