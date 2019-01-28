package repository

import (
	"forum/entitys"
	"github.com/jinzhu/gorm"
)

type LikeTopicRepo struct {
	db *gorm.DB
}

func NewLikeTopicRepo(db *gorm.DB) *LikeTopicRepo {
	return &LikeTopicRepo{db: db}
}

func (r *LikeTopicRepo) FindByUserIDAndTopicID(userID, topicID uint) (*entitys.LikeTopic, error) {
	var (
		likeTopic = new(entitys.LikeTopic)
		err       = r.db.Where("user_id = ? and topic_id = ?", userID, topicID).
				Order("created_at desc").First(likeTopic).Error
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return likeTopic, nil
		}
	}
	return likeTopic, err
}

func (r *LikeTopicRepo) CreateLike(likeTopic *entitys.LikeTopic) error {
	return r.db.Create(likeTopic).Error
}

func (r *LikeTopicRepo) CreateDislike(likeTopic *entitys.LikeTopic) error {
	return r.db.Create(likeTopic).Error
}

func (r *LikeTopicRepo) CancelLikeOrDislike(userID, topicID uint) (bool, error) {
	err := r.db.Where("user_id = ? and topic_id = ?", userID, topicID).Delete(&entitys.LikeTopic{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LikeTopicRepo) Update(topicLike *entitys.LikeTopic) error {
	return r.db.Save(topicLike).Error
}
