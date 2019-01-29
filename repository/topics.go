package repository

import (
	"forum/entitys"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
)

type TopicsRepo struct {
	db *gorm.DB
}

func NewTopicsRepo(db *gorm.DB) *TopicsRepo {
	return &TopicsRepo{db: db}
}

func (r *TopicsRepo) Create(topic *entitys.Topic) error {
	return r.db.Create(topic).Error
}

func (r *TopicsRepo) FindByID(id uint) (*entitys.Topic, error) {
	var topic entitys.Topic
	result := r.db.Where("id=?", id).Preload("Label").Preload("User").First(&topic)

	if err := result.Error; err != nil {
		return &topic, err
	}
	if result.RecordNotFound() == true {
		return &topic, errors.New("topic not found")
	}

	return &topic, nil
}

func (r *TopicsRepo) FindAll(limit, offset int) ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)

	result := r.db.Preload("Label").Preload("User").Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)

	result := r.db.Where("title != ''").Preload("Label", labelID).Preload("User").Order("top desc, last_reply_time desc, created_at desc").Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllByUserID(userID, limit, offset uint) ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)
	result := r.db.Preload("Label").Preload("User", userID).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllByUserIDOrderBy(userID uint, order string, limit, offset uint) ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)
	result := r.db.Where("user_id = ?", userID).Preload("Label").Preload("User").Order(order).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllNews() ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)
	result := r.db.Where("title <> ''").Preload("User").Preload("Label").Preload("LastReplyUser").
		Order("last_reply_time desc, created_at desc").
		Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindHots(limit int) ([]entitys.Topic, error) {
	var topics = make([]entitys.Topic, 0)
	result := r.db.Preload("User").Preload("Label").Order("replies_count desc, views_count desc").
		Limit(limit).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) GetTopicsCount(userID uint) map[string]int {
	var (
		result      = make(map[string]int)
		topicsCount = 0
		replyCount  = 0
		topicGoods  = 0
	)

	r.db.Model(&entitys.Topic{}).Where("user_id = ?", userID).Count(&topicsCount)
	r.db.Model(&entitys.Reply{}).Where("user_id = ?", userID).Count(&replyCount)
	r.db.Model(&entitys.Topic{}).Where("user_id = ? and good = ?", userID, true).Count(&topicGoods)

	result["TopicsCount"] = topicsCount
	result["ReplyCount"] = replyCount
	result["TopicGoods"] = topicGoods

	return result
}

func (r *TopicsRepo) Update(topic *entitys.Topic) error {
	return r.db.Save(topic).Error
}
