package repository

import (
	"fmt"
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var userRepo = NewUserRepo(datasource.InstanceGormMaster())

func TestUserRepo_Create(t *testing.T) {
	user := &entitys.User{
		Username: "test",
		Password: "testPassword",
		Email:    "test@forum.com",
		Url:      "",
		Avatar:   "",
	}
	err := userRepo.Create(user)
	if err != nil {
		t.Error(err)
	}
	userRepo.db.Unscoped().Delete(user)
}

func TestUserRepo_FindByID(t *testing.T) {
	user := &entitys.User{
		Username: "test",
		Password: "testPassword",
		Email:    "test@email.com",
	}
	err := userRepo.Create(user)
	if err != nil {
		t.Error(err)
	}

	_u, err := userRepo.FindByID(user.ID)
	if err != nil {
		fmt.Println(err)
		userRepo.db.Unscoped().Delete(user)
		t.Error(err)
	}
	if _u.ID != user.ID {
		t.Error("通过ID查找失败")
	}
	userRepo.db.Unscoped().Delete(user)
}

func TestUserRepo_FindByUsername(t *testing.T) {
	user := &entitys.User{
		Username: "test",
		Password: "testPassword",
		Email:    "test@email.com",
	}
	err := userRepo.Create(user)
	if err != nil {
		t.Error(err)
	}

	_u, err := userRepo.FindByUsername(user.Username)

	if err != nil {
		t.Error(err)
	}
	if _u.Username != user.Username {
		userRepo.db.Unscoped().Delete(user)
		t.Error("通过username查找失败")
	}
	userRepo.db.Unscoped().Delete(user)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	user := &entitys.User{
		Username: "test",
		Password: "testPassword",
		Email:    "test@email.com",
	}
	err := userRepo.Create(user)
	if err != nil {
		t.Error(err)
	}

	_u, err := userRepo.FindByEmail(user.Email)
	if err != nil {
		t.Error(err)
	}
	if _u.Email != user.Email {
		fmt.Println("_u.email:", _u.Email)
		fmt.Println("user.email:", user.Email)
		userRepo.db.Unscoped().Delete(user)
		t.Error("通过Email查找失败")
	}
	userRepo.db.Unscoped().Delete(user)
}

func TestUserRepo_Update(t *testing.T) {
	user := entitys.User{
		Username: "test",
		Password: "testPassword",
		Email:    "test@email.com",
	}
	err := userRepo.Create(&user)
	if err != nil {
		t.Error(err)
	}
	user.Email = "zzz@email.com"
	err = userRepo.Update(&user)
	if err != nil {
		t.Error("更新失败:", err)
	}

	_u, err := userRepo.FindByEmail(user.Email)
	if err != nil {
		userRepo.db.Unscoped().Delete(user)
		t.Error("查找失败:", err)
	}
	if _u.Email != user.Email {
		userRepo.db.Unscoped().Delete(user)
		t.Error("更新失败")
	}
	userRepo.db.Unscoped().Delete(user)
}
