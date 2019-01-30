package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
)

type MessageService interface {
	GetMessage(userID, mID uint) (*entitys.Message, error)
	GetAllMessagesByUser(userID, limit, offset uint) ([]entitys.Message, error)
	GetAllNotReadMessages(userID, limit, offset uint) ([]entitys.Message, error)

	ReadMessage(userID, mID uint) error
	DeleteMessage(userID, mID uint) error
	DeleteMessages(userID uint, mIDs []uint) error

	SendMessageToUser(userID, reID uint, title, content, url string) error
	SendMessagesToAllUser(title, content, url string) error
}

type messageService struct {
	repo *repository.MessageRepo
}

func (s *messageService) GetMessage(userID, mID uint) (*entitys.Message, error) {
	return s.repo.FindMessageByUser(userID, mID)
}

func (s *messageService) GetAllMessagesByUser(userID, limit, offset uint) ([]entitys.Message, error) {
	offset--
	return s.repo.FindFullMessageByUserID(userID, limit, offset)
}

func (s *messageService) GetAllNotReadMessages(userID, limit, offset uint) ([]entitys.Message, error) {
	offset--
	return s.repo.FindAllUnreadMessageByUserID(userID, limit, offset)
}

func (s *messageService) ReadMessage(userID, mID uint) error {
	return s.repo.ReadMessage(userID, mID)
}

func (s *messageService) DeleteMessage(userID, mID uint) error {
	return s.repo.DeleteMessage(userID, mID)
}

func (s *messageService) DeleteMessages(userID uint, mIDs []uint) error {
	return s.repo.DeleteMessages(userID, mIDs)
}

func (s *messageService) SendMessageToUser(userID, reID uint, title, content, url string) error {
	messageText := &entitys.MessageText{
		SenderID: userID,
		Title:    title,
		Content:  content,
		Url:      url,
	}
	message := &entitys.Message{
		ReceiverUserID: reID,
	}

	return s.repo.SendMessage(messageText, message)
}

func (s *messageService) SendMessagesToAllUser(title, content, url string) error {
	panic("implement me")
}

func NewMessageService() MessageService {
	return &messageService{
		repo: repository.NewMessageRepo(datasource.InstanceGormMaster()),
	}
}
