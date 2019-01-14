package repository

import (
	"fmt"
	"forum/entitys"
	"forum/pkg/business_errors"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entitys.Users) error {
	if err := r.db.Create(user).Error; err != nil {
		fmt.Printf("CreateUserError:%s", err)
		return err
	}
	return nil
}

func (r *UserRepo) FindByID(id uint) (*entitys.Users, error) {
	var user = new(entitys.Users)
	result := r.db.Where("id=?", id).First(user)
	err := result.Error
	return user, err
}

func (r *UserRepo) FindByUsername(username string) (*entitys.Users, error) {
	var user = new(entitys.Users)
	result := r.db.Where("username = ?", username).Find(&user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepo) FindByEmail(email string) (*entitys.Users, error) {
	var user = new(entitys.Users)
	result := r.db.Where("email=?", email).Find(user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepo) Update(users *entitys.Users) error {
	return r.db.Save(users).Error
}

func (r *UserRepo) Delete(user *entitys.Users) error {
	err := r.db.Unscoped().Delete(user).Error
	return err
}

func (r *UserRepo) DeleteByID(userID uint) error {
	var user = new(entitys.Users)
	user.ID = userID
	return r.db.Unscoped().Delete(user).Error
}
