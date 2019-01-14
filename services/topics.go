package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"time"
)

type TopicsService interface {
	Create(topics *entitys.Topics) error

	FindByID(topicID uint) (*entitys.Topics, error)
	FindAll(limit, offset int) ([]entitys.Topics, error)
	FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topics, error)
	FindAllByUserID(userID, limit, offset uint) ([]entitys.Topics, error)
	FindAllNewsTopics() ([]entitys.Topics, error)
	FindHots(limit int) []entitys.Topics

	AddTopicOnceViewCount(topic *entitys.Topics)
	ReplyTopicHandle(topic *entitys.Topics, replyUserID uint)
}

var (
	_labelServers = NewLabelService()
)

type topicsService struct {
	repo *repository.TopicsRepo
}

func (s *topicsService) Create(topics *entitys.Topics) error {
	label, err := _labelServers.FindByID(topics.LabelId)
	if err != nil {
		return err
	}
	_labelServers.PostTopicHandle(label)
	topics.LastReplyTime = time.Now()
	return s.repo.Create(topics)
}

func (s *topicsService) FindByID(topicID uint) (*entitys.Topics, error) {
	topic, err := s.repo.FindByID(topicID)
	s.AddTopicOnceViewCount(topic)
	return topic, err
}

func (s *topicsService) FindAll(limit, offset int) ([]entitys.Topics, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *topicsService) FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topics, error) {
	return s.repo.FindAllByLabelID(labelID, limit, offset)
}

func (s *topicsService) FindAllByUserID(userID, limit, offset uint) ([]entitys.Topics, error) {
	return s.repo.FindAllByLabelID(userID, limit, offset)
}

func (s *topicsService) FindAllNewsTopics() ([]entitys.Topics, error) {
	return s.repo.FindAllNews()
}

func (s *topicsService) FindHots(limit int) []entitys.Topics {
	t, _ := s.repo.FindHots(limit)
	return t
}

// AddTopicOnceViewCount
// Increase the number of topic readings
func (s *topicsService) AddTopicOnceViewCount(topic *entitys.Topics) {
	topic.ViewsCount += 1
	_ = s.repo.Update(topic)
}

// ReplyTopicHandle
// Update topic information
func (s *topicsService) ReplyTopicHandle(topic *entitys.Topics, replyUserID uint) {
	topic.LastReplyTime = time.Now()
	topic.RepliesCount += 1
	topic.LastReplyUserID = replyUserID
	_ = s.repo.Update(topic)
}

func NewTopicsService() TopicsService {
	return &topicsService{
		repo: repository.NewTopicsRepo(datasource.InstanceGormMaster()),
	}
}
