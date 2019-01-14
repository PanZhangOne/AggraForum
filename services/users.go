package services

import (
	"fmt"
	"forum/datasource"
	"forum/entitys"
	"forum/pkg/business_errors"
	"forum/repository"
	"forum/util/crypto"
)

type UsersService interface {
	Create(user *entitys.Users) error
	Login(username, password string) (*entitys.Users, error)

	FindByID(userID uint) (*entitys.Users, error)
	FindByIDs(ids []uint) []*entitys.Users

	FindByUsername(username string) (*entitys.Users, error)

	UpdateUserInfo(user *entitys.Users, columns []string) (*entitys.Users, error)

	CheckUsernameExist(username string) bool
	CheckEmailExist(email string) bool
}

type userService struct {
	repo *repository.UserRepo
}

func (s *userService) Create(user *entitys.Users) error {
	if s.CheckEmailExist(user.Username) {
		return business_errors.UsernameAlreadyExists
	}

	if s.CheckEmailExist(user.Email) {
		return business_errors.EmailAlreadyExists
	}

	if len(user.Password) < 8 {
		return business_errors.PasswordLessThanEightCharacters
	}

	pw, err := crypto.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pw
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string) (*entitys.Users, error) {
	user, err := s.repo.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	if user.ID < 0 {
		return nil, business_errors.UsernameNotExist
	}

	isCheck := crypto.CheckPassword(password, user.Password)
	if !isCheck {
		return nil, business_errors.PasswordError
	}
	return user, nil
}

func (s *userService) FindByID(userID uint) (*entitys.Users, error) {
	user, err := s.repo.FindByID(userID)
	return user, err
}

func (s *userService) FindByIDs(ids []uint) []*entitys.Users {
	return nil
}

func (s *userService) FindByUsername(username string) (*entitys.Users, error) {
	return s.repo.FindByUsername(username)
}

func (s *userService) UpdateUserInfo(user *entitys.Users, columns []string) (*entitys.Users, error) {
	return nil, nil
}

func (s *userService) CheckUsernameExist(username string) bool {
	user, err := s.repo.FindByUsername(username)
	fmt.Println("found user:", user)
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func (s *userService) CheckEmailExist(email string) bool {
	user, err := s.repo.FindByEmail(email)
	fmt.Println("found user:", user)
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func NewUserService() UsersService {
	return &userService{
		repo: repository.NewUserRepo(datasource.InstanceGormMaster()),
	}
}
