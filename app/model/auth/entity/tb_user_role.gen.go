// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameTbUserRole = "tb_user_role"

// TbUserRole mapped from table <tb_user_role>
type TbUserRole struct {
	ID     int  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UserID *int `gorm:"column:user_id;type:int" json:"user_id"` // 用户id
	RoleID *int `gorm:"column:role_id;type:int" json:"role_id"` // 角色id
}

// TableName TbUserRole's table name
func (*TbUserRole) TableName() string {
	return TableNameTbUserRole
}
