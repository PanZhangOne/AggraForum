package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"github.com/kataras/iris/core/errors"
)

type CollectTopicService interface {
	FindByUserID(userID, limit, offset uint) ([]entitys.CollectTopic, error)

	Collect(userID, topicID, labelID uint) error
	UnCollect(userID, topicID uint) error

	CheckCollectedTopic(userID, topicID uint) bool
}

type collectTopicService struct {
	repo *repository.CollectTopicRepo
}

func (s *collectTopicService) FindByUserID(userID, limit, offset uint) ([]entitys.CollectTopic, error) {
	offset--
	return s.repo.FindByUserID(userID, limit, offset)
}

func (s *collectTopicService) Collect(userID, topicID, labelID uint) error {
	if s.CheckCollectedTopic(userID, topicID) {
		return errors.New("该主题已收藏")
	}
	_, err := s.repo.Collect(userID, topicID, labelID)
	return err
}

func (s *collectTopicService) UnCollect(userID, topicID uint) error {
	return s.repo.UnCollect(userID, topicID)
}

func (s *collectTopicService) CheckCollectedTopic(userID, topicID uint) bool {
	c, _ := s.repo.FindByUserIDAndTopicID(userID, topicID)
	return c.ID > 0
}

func NewCollectTopicService() CollectTopicService {
	return &collectTopicService{
		repo: repository.NewCollectTopicRepo(datasource.InstanceGormMaster()),
	}
}
