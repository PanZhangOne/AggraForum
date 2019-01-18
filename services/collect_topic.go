package services

type CollectTopicService interface {
	Collect(userID, topicID uint) error
	UnCollect(userID, topicID uint) error
}
