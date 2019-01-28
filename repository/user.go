package repository

import (
	"fmt"
	"forum/entitys"
	"forum/util/business_errors"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entitys.User) error {
	if err := r.db.Create(user).Error; err != nil {
		fmt.Printf("CreateUserError:%s", err)
		return err
	}
	return nil
}

func (r *UserRepo) FindByID(id uint) (*entitys.User, error) {
	var user = new(entitys.User)
	result := r.db.Where("id=?", id).First(user)
	err := result.Error
	return user, err
}

func (r *UserRepo) FindByIDs(ids []uint) ([]entitys.User, error) {
	var users = make([]entitys.User, 0)
	err := r.db.Where("id in (?)", ids).Find(&users).Error
	for idx, _ := range users {
		users[idx].Password = ""
	}
	return users, err
}

func (r *UserRepo) FindByUsername(username string) (*entitys.User, error) {
	var user = new(entitys.User)
	result := r.db.Where("username = ?", username).Find(&user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepo) FindByEmail(email string) (*entitys.User, error) {
	var user = new(entitys.User)
	result := r.db.Where("email=?", email).Find(user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepo) FindAllUsers(limit, offset int) ([]entitys.User, error) {
	var users = make([]entitys.User, 0)
	err := r.db.Find(&users).Error
	for idx, _ := range users {
		users[idx].Password = ""
	}
	return users, err
}

func (r *UserRepo) Update(users *entitys.User) error {
	return r.db.Save(users).Error
}

func (r *UserRepo) Delete(user *entitys.User) error {
	err := r.db.Unscoped().Delete(user).Error
	return err
}

func (r *UserRepo) DeleteByID(userID uint) error {
	var user = new(entitys.User)
	user.ID = userID
	return r.db.Unscoped().Delete(user).Error
}
