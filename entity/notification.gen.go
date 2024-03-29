// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameNotification = "notification"

// Notification mapped from table <notification>
type Notification struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	Title     string    `gorm:"column:title" json:"title"`
	Body      string    `gorm:"column:body" json:"body"`
	Assets    string    `gorm:"column:assets" json:"assets"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
