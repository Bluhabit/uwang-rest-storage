// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameGroup = "group"

// Group mapped from table <group>
type Group struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	GroupOwner       string    `gorm:"column:group_owner" json:"group_owner"`
	GroupBanner      string    `gorm:"column:group_banner" json:"group_banner"`
	GroupName        string    `gorm:"column:group_name" json:"group_name"`
	GroupDescription string    `gorm:"column:group_description" json:"group_description"`
	GroupMemberCount int64     `gorm:"column:group_member_count" json:"group_member_count"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted          bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}
