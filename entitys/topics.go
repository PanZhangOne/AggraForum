package entitys

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Topic struct {
	LabelId uint   `gorm:"not null;uint;foreignkey"`
	UserId  uint   `gorm:"not null;uint;foreignkey"`
	Title   string `gorm:"not null;varchar(128)"`
	Content string `gorm:"text"`

	ViewsCount   uint `gorm:"not null;uint"`
	RepliesCount uint `gorm:"not null;uint"`

	LastReplyUserID uint `gorm:"uint;foreignkey"`
	LastReplyTime   time.Time

	Like    int `gorm:"default:0"`
	Dislike int `gorm:"default:0"`

	Top  bool `gorm:"boolean;default:0"`
	Good bool `gorm:"boolean;default:0"`
	Lock bool `gorm:"boolean;default:0"`

	Label         Labels `gorm:"foreignkey:label_id"`
	User          User   `gorm:"foreignkey:user_id"`
	LastReplyUser User   `gorm:"foreignkey:last_reply_user_id"`
	gorm.Model
}
