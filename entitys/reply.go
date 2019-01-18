package entitys

import (
	"html/template"
	"time"
)

type Reply struct {
	ID       uint `gorm:"primary_key;AUTO_INCREMENT"`
	TopicID  uint `gorm:"not null;foreignkey"`
	UserID   uint `gorm:"not null;foreignkey"`
	ParentID uint
	Floor    int    `gorm:"not null;default:1"`
	Content  string `gorm:"type:text;not null"`

	DeviceInfo string
	Thanks     uint `gorm:"default:0"`

	CreatedAt time.Time

	User        User          `gorm:"foreignkey:user_id"`
	Topic       Topic         `gorm:"foreignkey:topic_id"`
	ContentHtml template.HTML `gorm:"-"`
	ParentReply struct {
		Floor int
	} `gorm:"-"`
}
