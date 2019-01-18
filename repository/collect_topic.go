package repository

import (
	"forum/entitys"
	"github.com/jinzhu/gorm"
)

type CollectTopicRepo struct {
	db *gorm.DB
}

func NewCollectTopicRepo(db *gorm.DB) *CollectTopicRepo {
	return &CollectTopicRepo{db: db}
}

func (r *CollectTopicRepo) Collect(userID, topicID uint) (*entitys.CollectTopic, error) {
	var collectTopic = new(entitys.CollectTopic)
	collectTopic.UserID = userID
	collectTopic.TopicID = topicID
	err := r.db.Create(collectTopic).Error
	return collectTopic, err
}

func (r *CollectTopicRepo) UnCollect(userID, topicID uint) error {
	var collectTopic = new(entitys.CollectTopic)
	collectTopic.UserID = userID
	collectTopic.TopicID = topicID
	err := r.db.Delete(collectTopic).Error
	return err
}
