package entitys

import "time"

type Users struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `gorm:"unique;not null;varchar(50)"`
	Password string `gorm:"not null;varchar(128)"`
	Email    string `gorm:"not null;unique;varchar(128)"`

	Url    string `gorm:"varchar(255)"`
	Avatar string `gorm:"varchar(255)"`

	Score             int  `gorm:"int;default:0;"`
	TopicCount        int  `gorm:"int;default:0"`
	ReplyCount        int  `gorm:"int;default:0"`
	FollowerCount     int  `gorm:"int;default:0"`
	FollowingCount    int  `gorm:"int;default:0"`
	CollectTagCount   int  `gorm:"int;default:0"`
	CollectTopicCount int  `gorm:"int;default:0"`
	Lock              bool `gorm:"boolean;default:0"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

//func (u *Users) BeforeCreate(err error) {
//	uid := uuid.NewV4().String()
//	u.UserID = uid
//	return
//}
