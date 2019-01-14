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

func (r *TopicsRepo) Create(topic *entitys.Topics) error {
	return r.db.Create(topic).Error
}

func (r *TopicsRepo) FindByID(id uint) (*entitys.Topics, error) {
	var topic entitys.Topics
	result := r.db.Where("id=?", id).Preload("Label").Preload("User").First(&topic)

	if err := result.Error; err != nil {
		return &topic, err
	}
	if result.RecordNotFound() == true {
		return &topic, errors.New("topic not found")
	}

	return &topic, nil
}

func (r *TopicsRepo) FindAll(limit, offset int) ([]entitys.Topics, error) {
	var topics = make([]entitys.Topics, 0)

	result := r.db.Preload("Label").Preload("User", ).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllByLabelID(labelID, limit, offset uint) ([]entitys.Topics, error) {
	var topics = make([]entitys.Topics, 0)

	result := r.db.Preload("Label", labelID).Preload("User").Order("last_reply_time desc, created_at desc").Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllByUserID(userID, limit, offset uint) ([]entitys.Topics, error) {
	var topics = make([]entitys.Topics, 0)
	result := r.db.Preload("Label").Preload("User", userID).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindAllNews() ([]entitys.Topics, error) {
	var topics = make([]entitys.Topics, 0)
	result := r.db.Preload("User").Preload("Label").Preload("LastReplyUser").Order("last_reply_time desc, created_at desc").
		Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) FindHots(limit int) ([]entitys.Topics, error) {
	var topics = make([]entitys.Topics, 0)
	result := r.db.Preload("User").Preload("Label").Order("replies_count desc, views_count desc").
		Limit(limit).Find(&topics)
	err := result.Error
	return topics, err
}

func (r *TopicsRepo) Update(topic *entitys.Topics) error {
	return r.db.Save(topic).Error
}
