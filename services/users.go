package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
	"forum/util/business_errors"
	"forum/util/crypto"
)

type UsersService interface {
	Create(user *entitys.User) error
	Login(username, password string) (*entitys.User, error)

	FindByID(userID uint) (*entitys.User, error)
	FindByIDs(ids []uint) []entitys.User

	FindByUsername(username string) (*entitys.User, error)
	FindAllUsers(limit, offset int) ([]entitys.User, error)

	CheckUsernameExist(username string) bool
	CheckEmailExist(email string) bool
	CheckUserIsLockByID(userID uint) bool

	// Actions
	LockUser(user *entitys.User)
	UnLockUser(user *entitys.User)
}

type userService struct {
	repo *repository.UserRepo
}

func (s *userService) Create(user *entitys.User) error {
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

func (s *userService) Login(username, password string) (*entitys.User, error) {
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

func (s *userService) FindByID(userID uint) (*entitys.User, error) {
	user, err := s.repo.FindByID(userID)
	return user, err
}

func (s *userService) FindByIDs(ids []uint) []entitys.User {
	users, _ := s.repo.FindByIDs(ids)
	return users
}

func (s *userService) FindByUsername(username string) (*entitys.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *userService) FindAllUsers(limit, offset int) ([]entitys.User, error) {
	return s.repo.FindAllUsers(limit, offset)
}

func (s *userService) CheckUsernameExist(username string) bool {
	user, err := s.repo.FindByUsername(username)
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
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func (s *userService) LockUser(user *entitys.User) {
	user.Lock = true
	_ = s.repo.Update(user)
}

func (s *userService) UnLockUser(user *entitys.User) {
	user.Lock = false
	_ = s.repo.Update(user)
}

func (s *userService) CheckUserIsLockByID(userID uint) bool {
	user, _ := s.FindByID(userID)
	return user.Lock
}

func NewUserService() UsersService {
	return &userService{
		repo: repository.NewUserRepo(datasource.InstanceGormMaster()),
	}
}
