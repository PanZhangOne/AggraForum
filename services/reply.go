package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
)

var _topicServices = NewTopicsService()

type RepliesService interface {
	Reply(reply *entitys.Reply) error
	FindRepliesByTopicID(topicID uint) []entitys.Reply
}

type repliesService struct {
	repo *repository.ReplyRepo
}

func (s *repliesService) Reply(reply *entitys.Reply) error {
	topic, err := _topicServices.FindByID(reply.TopicID)
	reply.Floor = int(topic.RepliesCount + 1)
	if err != nil {
		return err
	}
	_topicServices.ReplyTopicHandle(topic, reply.UserID)
	return s.repo.Create(reply)
}

func (s *repliesService) FindRepliesByTopicID(topicID uint) []entitys.Reply {
	replies, _ := s.repo.FindByTopicID(topicID)
	return replies
}

// ThanksReply 感谢回帖
func (s *repliesService) ThanksReply(replyID uint) error {
	reply, err := s.repo.FindByID(replyID)
	if err != nil {
		return err
	}
	reply.Thanks += 1
	err = s.repo.Update(reply)
	if err != nil {
		return err
	}
	return nil
}

func NewRepliesService() RepliesService {
	return &repliesService{repo: repository.NewReplyRepo(datasource.InstanceGormMaster())}
}
