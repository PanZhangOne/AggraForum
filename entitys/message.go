package entitys

import "time"

type Message struct {
	ID             uint `gorm:"primary_key;AUTO_INCREMENT"`
	ReceiverUserID uint `gorm:"not null;foreignkey"`
	MessageID      uint `gorm:"not null;foreignkey"`
	Status         int  `gorm:"default:'0';comment:'0未读,1已读'"`

	MessageText MessageText `gorm:"foreignkey:message_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageText struct {
	ID       uint `gorm:"primary_key;AUTO_INCREMENT"`
	SenderID uint `gorm:"not null;foreignkey"`

	Title   string `gorm:"varchar(50);not null;"`
	Content string `gorm:"varchar(255);not null;"`
	Url     string

	GroupID      uint `gorm:"comment:'用户组'"`
	MessageType  int  `gorm:"comment:'1 private, 2 public, 3 global'"`
	BusinessType int

	CreatedAt time.Time
	UpdatedAt time.Time
}
