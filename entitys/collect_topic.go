package entitys

import (
	"time"
)

type CollectTopic struct {
	ID      uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID  uint `gorm:"not null;foreignkey"`
	TopicID uint `gorm:"not null;foreignkey"`

	Topic     Topic `gorm:"foreignkey:topic_id"`
	CreatedAt time.Time
}
