package entitys

import "time"

type LikeTopic struct {
	ID      uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID  uint `gorm:"not null;foreignkey"`
	TopicID uint `gorm:"not null;foreignkey"`
	LabelID uint `gorm:"not null;foreignkey"`

	Like    bool
	Dislike bool

	User  User  `gorm:"foreignkey:user_id"`
	Topic Topic `gorm:"foreignkey:topic_id"`
	Label Label `gorm:"foreignkey:label_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
