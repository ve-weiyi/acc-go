// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameUserAuth = "user_auth"

// UserAuth mapped from table <user_auth>
type UserAuth struct {
	ID       int    `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);primaryKey" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
}

// TableName UserAuth's table name
func (*UserAuth) TableName() string {
	return TableNameUserAuth
}