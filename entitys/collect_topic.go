package entitys

import (
	"time"
)

type CollectTopic struct {
	UserID  uint `gorm:"not null;foreignkey"`
	TopicID uint `gorm:"not null;foreignkey"`

	Topic     Topic
	CreatedAt time.Time
}
