package repository

import (
	"forum/entitys"
	"github.com/jinzhu/gorm"
)

type ReplyRepo struct {
	db *gorm.DB
}

func NewReplyRepo(db *gorm.DB) *ReplyRepo {
	return &ReplyRepo{db: db}
}

func (r *ReplyRepo) Create(reply *entitys.Reply) error {
	return r.db.Create(reply).Error
}

func (r *ReplyRepo) FindByID(replyID uint) (*entitys.Reply, error) {
	var reply = new(entitys.Reply)
	reply.ID = replyID
	err := r.db.First(reply).Error
	return reply, err
}

func (r *ReplyRepo) FindByIDs(ids []uint) ([]entitys.Reply, error) {
	var replies = make([]entitys.Reply, 0)
	err := r.db.Where("id in (?)", ids).Find(&replies).Error
	return replies, err
}

func (r *ReplyRepo) FindByUserID(userID, limit, offset uint) ([]entitys.Reply, error) {
	var replies = make([]entitys.Reply, 0)
	err := r.db.Where("user_id = ?", userID).Preload("Topic").Limit(limit).Offset(offset).Find(&replies).Error
	return replies, err
}

func (r *ReplyRepo) FindByTopicID(topicID uint) ([]entitys.Reply, error) {
	var replies = make([]entitys.Reply, 0)
	result := r.db.Where("topic_id = ?", topicID).Preload("User").Find(&replies)
	err := result.Error
	var replyIds = make([]uint, 0)
	for _, reply := range replies {
		if reply.ParentID > 0 {
			replyIds = append(replyIds, reply.ParentID)
		}
	}
	parentReplies, _ := r.FindByIDs(replyIds)
	for idx, reply := range replies {
		if reply.ParentID > 0 {
			for _, parentReply := range parentReplies {
				if reply.ParentID == parentReply.ID {
					replies[idx].ParentReply.Floor = parentReply.Floor
				}
			}
		}
	}
	return replies, err
}

func (r *ReplyRepo) Update(reply *entitys.Reply) error {
	return r.db.Save(reply).Error
}
