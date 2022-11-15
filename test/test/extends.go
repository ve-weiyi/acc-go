package test

import "acc/app/model/entity"

type UserBean struct {
	Auth    entity.UserAuth
	Details entity.UserInfo
	Roles   []string
}
