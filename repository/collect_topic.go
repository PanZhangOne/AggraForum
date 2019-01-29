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

func (r *CollectTopicRepo) FindByID(id uint) (*entitys.CollectTopic, error) {
	var collectTopic = new(entitys.CollectTopic)
	err := r.db.Where("id = ?", id).First(collectTopic).Error
	return collectTopic, err
}

func (r *CollectTopicRepo) FindByUserID(userID, limit, offset uint) ([]entitys.CollectTopic, error) {
	var (
		topicCollects = make([]entitys.CollectTopic, 0)
		err           = r.db.Where("user_id = ?", userID).Preload("Topic").Limit(limit).Offset(offset).
				Order("created_at desc").Find(&topicCollects).Error
	)
	return topicCollects, err
}

func (r *CollectTopicRepo) FindByTopicID(id uint) (*entitys.CollectTopic, error) {
	var collectTopic = new(entitys.CollectTopic)
	err := r.db.Where("topic_id = ?", id).First(collectTopic).Error
	return collectTopic, err
}

func (r *CollectTopicRepo) FindByUserIDAndTopicID(userID, topicID uint) (*entitys.CollectTopic, error) {
	var collectTopic = new(entitys.CollectTopic)
	err := r.db.Where("topic_id = ? and user_id = ?", topicID, userID).First(collectTopic).Error
	return collectTopic, err
}

func (r *CollectTopicRepo) Collect(userID, topicID, LabelID uint) (*entitys.CollectTopic, error) {
	var collectTopic = new(entitys.CollectTopic)
	collectTopic.UserID = userID
	collectTopic.TopicID = topicID
	collectTopic.LabelID = LabelID
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
