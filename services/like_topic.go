package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"github.com/kataras/iris/core/errors"
)

type LikeTopicService interface {
	Like(userID, topicID uint) (bool, error)
	Dislike(userID, topicID uint) (bool, error)

	CancelLike(userID, topicID uint) (bool, error)
	CancelDislike(userID, topicID uint) (bool, error)

	FindTopicIsLikeOrDislike(userID, topicID uint) (*entitys.LikeTopic, error)
}

type likeTopicService struct {
	repo *repository.LikeTopicRepo
}

var (
	topicService = NewTopicsService()
)

func (s *likeTopicService) Like(userID, topicID uint) (bool, error) {
	var (
		topic, _  = topicService.FindByID(topicID)
		topicLike = new(entitys.LikeTopic)
	)
	if topic.ID <= 0 {
		return false, errors.New("该主题不存在")
	}

	topicLike.UserID = userID
	topicLike.TopicID = topicID
	topicLike.LabelID = topic.LabelId
	topicLike.Like = true

	record, err := s.repo.FindByUserIDAndTopicID(userID, topicID)

	if err != nil {
		return false, err
	}

	if record.ID >= 1 {
		if record.Like {
			return false, errors.New("该主题已经设置了喜欢")
		}
		if record.Dislike {
			return false, errors.New("该主题已经设置了不喜欢")
		}
	}

	err = s.repo.CreateLike(topicLike)
	if err != nil {
		return false, err
	}
	topicService.AddLikeCount(topic)
	return true, nil
}

func (s *likeTopicService) Dislike(userID, topicID uint) (bool, error) {
	var (
		topic, _  = topicService.FindByID(topicID)
		topicLike = new(entitys.LikeTopic)
	)

	if topic.ID <= 0 {
		return false, errors.New("该主题不存在")
	}

	record, err := s.repo.FindByUserIDAndTopicID(userID, topicID)
	if err != nil {
		return false, err
	}

	if record.ID >= 1 {
		if record.Like {
			return false, errors.New("该主题已经设置了喜欢")
		}
		if record.Dislike {
			return false, errors.New("该主题已经设置了不喜欢")
		}
	}

	topicLike.UserID = userID
	topicLike.TopicID = topicID
	topicLike.LabelID = topic.LabelId
	topicLike.Dislike = true

	err = s.repo.CreateDislike(topicLike)

	if err != nil {
		return false, err
	}
	topicService.AddDislikeCount(topic)
	return true, nil
}

func (s *likeTopicService) CancelLike(userID, topicID uint) (bool, error) {
	var (
		topic, _ = topicService.FindByID(topicID)
	)

	if topic.ID <= 0 {
		return false, errors.New("该主题不存在")
	}
	ok, err := s.repo.CancelLikeOrDislike(userID, topicID)
	if err != nil {
		return ok, err
	}
	topicService.ReduceLikeCount(topic)
	return true, nil
}

func (s *likeTopicService) CancelDislike(userID, topicID uint) (bool, error) {
	var (
		topic, _ = topicService.FindByID(topicID)
	)

	if topic.ID <= 0 {
		return false, errors.New("该主题不存在")
	}
	ok, err := s.repo.CancelLikeOrDislike(userID, topicID)
	if err != nil {
		return ok, err
	}
	topicService.ReduceDislikeCount(topic)
	return true, nil
}

func (s *likeTopicService) FindTopicIsLikeOrDislike(userID, topicID uint) (*entitys.LikeTopic, error) {
	return s.repo.FindByUserIDAndTopicID(userID, topicID)
}

func NewLikeTopicService() LikeTopicService {
	return &likeTopicService{
		repo: repository.NewLikeTopicRepo(datasource.InstanceGormMaster()),
	}
}
