package entitys

import "time"

type Label struct {
	ID            uint   `gorm:"primary_key;AUTO_INCREMENT"`
	LabelName     string `gorm:"varchar(32);not null;unique"`
	LabelDesc     string `gorm:"varchar(255);"`
	TopicsCount   int64  `gorm:"default:0"`
	CreatedUserID uint   `gorm:"uint"`
	Recommend     bool   `gorm:"boolean;default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
