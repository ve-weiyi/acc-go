// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameTbUserAuth = "tb_user_auth"

// TbUserAuth mapped from table <tb_user_auth>
type TbUserAuth struct {
	ID            int        `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UserInfoID    int        `gorm:"column:user_info_id;type:int;not null" json:"user_info_id"`                                                 // 用户信息id
	Username      string     `gorm:"column:username;type:varchar(50);not null;uniqueIndex:username,priority:1" json:"username"`                 // 用户名
	Password      string     `gorm:"column:password;type:varchar(100);not null" json:"password"`                                                // 密码
	LoginType     int        `gorm:"column:login_type;type:tinyint(1);not null" json:"login_type"`                                              // 登录类型
	IPAddress     *string    `gorm:"column:ip_address;type:varchar(50)" json:"ip_address"`                                                      // 用户登录ip
	IPSource      *string    `gorm:"column:ip_source;type:varchar(50)" json:"ip_source"`                                                        // ip来源
	CreateTime    time.Time  `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time" example:"2022-11-16T16:00:00.000Z"` // 创建时间
	UpdateTime    *time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time" example:"2022-11-16T16:00:00.000Z"` // 更新时间
	LastLoginTime *time.Time `gorm:"column:last_login_time;type:datetime" json:"last_login_time" example:"2022-11-16T16:00:00.000Z"`            // 上次登录时间
}

// TableName TbUserAuth's table name
func (*TbUserAuth) TableName() string {
	return TableNameTbUserAuth
}
