package services

import (
	"forum/datasource"
	"forum/repository"
	"github.com/kataras/iris/core/errors"
)

type CollectTopicService interface {
	Collect(userID, topicID uint) error
	UnCollect(userID, topicID uint) error
}

type collectTopicService struct {
	repo *repository.CollectTopicRepo
}

func (s *collectTopicService) Collect(userID, topicID uint) error {
	if s.CheckCollectedTopic(topicID) {
		return errors.New("该主题已收藏")
	}
	_, err := s.repo.Collect(userID, topicID)
	return err
}

func (s *collectTopicService) UnCollect(userID, topicID uint) error {
	return s.repo.UnCollect(userID, topicID)
}

func (s *collectTopicService) CheckCollectedTopic(topicID uint) bool {
	c, _ := s.repo.FindByTopicID(topicID)
	return c.ID > 0
}

func NewCollectTopicService() CollectTopicService {
	return &collectTopicService{
		repo: repository.NewCollectTopicRepo(datasource.InstanceGormMaster()),
	}
}
