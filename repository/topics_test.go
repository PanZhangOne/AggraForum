package repository

import (
	"fmt"
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var topicRepo = NewTopicsRepo(datasource.InstanceGormMaster())

func TestTopicsRepo_Create(t *testing.T) {
	var topic = new(entitys.Topic)
	topic.UserId = 7
	topic.LabelId = 1
	topic.Title = "test"
	topic.Content = "test"
	err := topicRepo.db.Create(topic).Error
	if err != nil {
		t.Error(err)
	}
}

func TestTopicsRepo_FindByID(t *testing.T) {
	res, err := topicRepo.FindByID(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("topic_title:", res.Title)
	fmt.Println("username:", res.User.Username)
	fmt.Println("password:", res.User.Password)
	fmt.Println("labelName:", res.Label.LabelName)
}

func TestTopicsRepo_FindAll(t *testing.T) {
	res, err := topicRepo.FindAll(50, 0)
	if err != nil {
		t.Error(err)
	}
	if len(res) <= 0 {
		t.Error("查找失败")
	}
	fmt.Println(res)
}

func TestTopicsRepo_FindAllByLabelID(t *testing.T) {
	res, err := topicRepo.FindAllByLabelID(1, 50, 0)
	if err != nil {
		t.Error(err)
	}
	if len(res) <= 0 {
		t.Error("查找失败")
	}
	fmt.Println(res)
}

func TestTopicsRepo_FindAllByUserID(t *testing.T) {
	res, err := topicRepo.FindAllByUserID(7, 50, 0)
	if err != nil {
		t.Error(err)
	}
	if len(res) <= 0 {
		t.Error("查找失败")
	}
	fmt.Println(res)
}
