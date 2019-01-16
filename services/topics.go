package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"time"
)

type TopicsService interface {
	Create(topics *entitys.Topic) error

	FindByID(topicID uint) (*entitys.Topic, error)
	FindAll(limit, offset int) ([]entitys.Topic, error)
	FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topic, error)
	FindAllByUserID(userID, limit, offset uint) ([]entitys.Topic, error)
	FindAllNewsTopics() ([]entitys.Topic, error)
	FindHots(limit int) []entitys.Topic

	AddTopicOnceViewCount(topic *entitys.Topic)
	ReplyTopicHandle(topic *entitys.Topic, replyUserID uint)
}

var (
	_labelServers = NewLabelService()
)

type topicsService struct {
	repo *repository.TopicsRepo
}

func (s *topicsService) Create(topics *entitys.Topic) error {
	label, err := _labelServers.FindByID(topics.LabelId)
	if err != nil {
		return err
	}
	_labelServers.PostTopicHandle(label)
	topics.LastReplyTime = time.Now()
	return s.repo.Create(topics)
}

func (s *topicsService) FindByID(topicID uint) (*entitys.Topic, error) {
	topic, err := s.repo.FindByID(topicID)
	s.AddTopicOnceViewCount(topic)
	return topic, err
}

func (s *topicsService) FindAll(limit, offset int) ([]entitys.Topic, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *topicsService) FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topic, error) {
	return s.repo.FindAllByLabelID(labelID, limit, offset)
}

func (s *topicsService) FindAllByUserID(userID, limit, offset uint) ([]entitys.Topic, error) {
	return s.repo.FindAllByLabelID(userID, limit, offset)
}

func (s *topicsService) FindAllNewsTopics() ([]entitys.Topic, error) {
	return s.repo.FindAllNews()
}

func (s *topicsService) FindHots(limit int) []entitys.Topic {
	t, _ := s.repo.FindHots(limit)
	return t
}

// AddTopicOnceViewCount
// Increase the number of topic readings
func (s *topicsService) AddTopicOnceViewCount(topic *entitys.Topic) {
	topic.ViewsCount += 1
	_ = s.repo.Update(topic)
}

// ReplyTopicHandle
// Update topic information
func (s *topicsService) ReplyTopicHandle(topic *entitys.Topic, replyUserID uint) {
	topic.LastReplyTime = time.Now()
	topic.RepliesCount += 1
	topic.LastReplyUserID = replyUserID
	_ = s.repo.Update(topic)
}

func (s *topicsService) TopTopicHandle(topic *entitys.Topic) {}

func NewTopicsService() TopicsService {
	return &topicsService{
		repo: repository.NewTopicsRepo(datasource.InstanceGormMaster()),
	}
}
