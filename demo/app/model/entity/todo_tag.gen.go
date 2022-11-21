// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameTodoTag = "todo_tag"

// TodoTag mapped from table <todo_tag>
type TodoTag struct {
	ID     int `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	TodoID int `gorm:"column:todo_id;type:int;not null;uniqueIndex:todo_id,priority:1" json:"todo_id"`
	TagID  int `gorm:"column:tag_id;type:int;not null;uniqueIndex:todo_id,priority:2" json:"tag_id"`
}

// TableName TodoTag's table name
func (*TodoTag) TableName() string {
	return TableNameTodoTag
}