package entitys

import (
	"time"
)

type CollectTopic struct {
	ID      uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID  uint `gorm:"not null;foreignkey"`
	TopicID uint `gorm:"not null;foreignkey"`
	LabelID uint `gorm:"not null;foreignkey"`

	Topic     Topic `gorm:"foreignkey:topic_id"`
	User      User  `gorm:"foreignkey:user_id"`
	Label     Label `gorm:"foreignkey:label_id"`
	CreatedAt time.Time
}
