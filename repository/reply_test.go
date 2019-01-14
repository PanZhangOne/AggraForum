package repository

import (
	"fmt"
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var replyRepo = NewReplyRepo(datasource.InstanceGormMaster())

func TestReplyRepo_Create(t *testing.T) {
	var reply = new(entitys.Reply)
	reply.UserID = 1
	reply.Content = "test"
	reply.TopicID = 6
	err := replyRepo.Create(reply)
	if err != nil {
		t.Error(err)
	}
}

func TestReplyRepo_FindByTopicID(t *testing.T) {
	var id = 6
	res, err := replyRepo.FindByTopicID(uint(id))
	if err != nil {
		t.Error(err)
	}
	if len(res) <= 0 {
		t.Error("查找失败")
	}
	fmt.Println(res)
}
