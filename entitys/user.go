package entitys

import "time"

type User struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `gorm:"unique;not null;varchar(50)"`
	Password string `gorm:"not null;varchar(128)"`
	Email    string `gorm:"not null;unique;varchar(128)"`

	Url    string `gorm:"varchar(255)"`
	Avatar string `gorm:"varchar(255)"`

	Score int  `gorm:"int;default:0;"`
	Lock  bool `gorm:"boolean;default:0"`

	TopicCount        int `gorm:"-"`
	ReplyCount        int `gorm:"-"`
	FollowerCount     int `gorm:"-"`
	FollowingCount    int `gorm:"-"`
	CollectTagCount   int `gorm:"-"`
	CollectTopicCount int `gorm:"-"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
