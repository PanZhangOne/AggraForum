package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"time"
)

type TopicsService interface {
	Create(topics *entitys.Topic) error

	// Finds
	FindByID(topicID uint) (*entitys.Topic, error)
	FindAll(limit, offset int) ([]entitys.Topic, error)
	FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topic, error)
	FindAllByUserID(userID, limit, offset uint) ([]entitys.Topic, error)
	FindAllNewTopicByUserID(userID, limit, offset uint) ([]entitys.Topic, error)
	FindAllNewsTopics() ([]entitys.Topic, error)
	FindHots(limit int) []entitys.Topic

	// Get
	GetTopicCount(userID uint) map[string]int

	// Actions
	AddTopicOnceViewCount(topic *entitys.Topic)
	ReplyTopicHandle(topic *entitys.Topic, replyUserID uint)
	ReduceLikeCount(topic *entitys.Topic)
	ReduceDislikeCount(topic *entitys.Topic)

	AddTopForTopic(topic *entitys.Topic)
	RemoveTopForTopic(topic *entitys.Topic)
	MoveTopicToLabel(topic *entitys.Topic, labelID uint)
	AddLikeCount(topic *entitys.Topic)
	AddDislikeCount(topic *entitys.Topic)
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

func (s *topicsService) FindAllNewTopicByUserID(userID, limit, offset uint) ([]entitys.Topic, error) {
	return s.repo.FindAllByUserIDOrderBy(userID, "created_at desc", limit, offset)
}

func (s *topicsService) FindAllNewsTopics() ([]entitys.Topic, error) {
	return s.repo.FindAllNews()
}

func (s *topicsService) FindHots(limit int) []entitys.Topic {
	t, _ := s.repo.FindHots(limit)
	return t
}

// GetTopicCount 获取贴数量
func (s *topicsService) GetTopicCount(userID uint) map[string]int {
	return s.repo.GetTopicsCount(userID)
}

//
//func (s *topicsService) Get

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

// AddTopForTopic 帖子置顶
func (s *topicsService) AddTopForTopic(topic *entitys.Topic) {
	topic.Top = true
	_ = s.repo.Update(topic)
}

// RemoveTopForTopic 帖子取消置顶
func (s *topicsService) RemoveTopForTopic(topic *entitys.Topic) {
	topic.Top = false
	_ = s.repo.Update(topic)
}

// MoveTopicToLabel 移动帖子到某个标签
func (s *topicsService) MoveTopicToLabel(topic *entitys.Topic, labelID uint) {
	topic.LabelId = labelID
	_ = s.repo.Update(topic)
}

// AddLikeCount 增加喜欢
func (s *topicsService) AddLikeCount(topic *entitys.Topic) {
	topic.LikeCount += 1
	_ = s.repo.Update(topic)
}

// AddDislikeCount 增加不喜欢
func (s *topicsService) AddDislikeCount(topic *entitys.Topic) {
	topic.DislikeCount += 1
	_ = s.repo.Update(topic)
}

// ReduceLikeCount 减少喜欢次数
func (s *topicsService) ReduceLikeCount(topic *entitys.Topic) {
	topic.LikeCount -= 1
	_ = s.repo.Update(topic)
}

// ReduceDislikeCount 减少不喜欢次数
func (s *topicsService) ReduceDislikeCount(topic *entitys.Topic) {
	topic.DislikeCount -= 1
	_ = s.repo.Update(topic)
}

func NewTopicsService() TopicsService {
	return &topicsService{
		repo: repository.NewTopicsRepo(datasource.InstanceGormMaster()),
	}
}
