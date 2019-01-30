package repository

import (
	"forum/entitys"
	"forum/util/business_types/message_status"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) FindMessageByUser(userID, mID uint) (*entitys.Message, error) {
	var message = new(entitys.Message)

	err := r.db.Where("receiver_user_id = ? and id = ?", userID, mID).Preload("MessageText").First(message).Error
	return message, err
}

func (r *MessageRepo) FindMessagesByUser(userID uint, mIDs []uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("receiver_user_id = ? and id in (?)", userID, mIDs).Preload("MessageText").Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindMessageByID(id uint) (*entitys.Message, error) {
	var message = new(entitys.Message)

	err := r.db.Where("id = ?", id).First(message).Error
	return message, err
}

func (r *MessageRepo) FindMessageByIDs(ids []uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("id in (?)", ids).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindFullMessageByID(id uint) (*entitys.Message, error) {
	var message = new(entitys.Message)

	err := r.db.Where("id = ?", id).Preload("MessageText").First(message).Error
	return message, err
}

func (r *MessageRepo) FindFullMessageByIDs(ids []uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("id in (?)", ids).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindMessageByUserID(userID, limit, offset uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("receiver_user_id = ?", userID).Order("created_at desc").Limit(limit).
		Offset(offset).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindFullMessageByUserID(userID, limit, offset uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("receiver_user_id = ?", userID).Preload("MessageText").Order("created_at desc").Limit(limit).
		Offset(offset).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindAllUnreadMessageByUserID(userID, limit, offset uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("receiver_user_id = ? and status = ?", userID, message_status.MessageUnread).
		Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) FindAllReadMessageByUserID(userID, limit, offset uint) ([]entitys.Message, error) {
	var messages = make([]entitys.Message, 0)

	err := r.db.Where("receiver_user_id = ? and status = ?", userID, message_status.MessageRead).
		Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) SendMessage(messageText *entitys.MessageText, message *entitys.Message) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(messageText).Error; err != nil {
		tx.Rollback()
		return err
	}

	message.MessageID = messageText.ID

	if err := tx.Create(message).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *MessageRepo) ReadMessage(userID, mID uint) error {
	var message = new(entitys.Message)
	err := r.db.Where("receiver_user_id = ? and message_id = ?", userID, mID).
		First(message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("未找到该信息")
		}
		return err
	}

	if message.Status != message_status.MessageUnread {
		return errors.New("该消息已读!")
	}
	message.Status = message_status.MessageRead
	return r.db.Save(message).Error
}

func (r *MessageRepo) DeleteMessage(userID, mID uint) error {
	var message = new(entitys.Message)
	err := r.db.Where("receiver_user_id = ? and message_id = ?", userID, mID).
		First(message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("未找到该信息")
		}
		return err
	}

	return r.db.Delete(message).Error
}

func (r *MessageRepo) DeleteMessages(userID uint, mIDs []uint) error {
	return r.db.Where("receiver_user_id = ? and id in (?)", userID, mIDs).
		Delete(&entitys.Message{}).Error
}
