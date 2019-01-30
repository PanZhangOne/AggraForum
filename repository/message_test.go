package repository

import (
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var messageRepo = NewMessageRepo(datasource.InstanceGormMaster())

func TestMessageRepo_SendMessage(t *testing.T) {
	var messageText = new(entitys.MessageText)
	messageText.Title = "Test"
	messageText.Content = "Test content"
	messageText.SenderID = 1
	messageText.MessageType = 1
	var message = new(entitys.Message)

	message.ReceiverUserID = 2
	message.MessageID = messageText.ID

	err := messageRepo.SendMessage(messageText, message)

	if err != nil {
		t.Error(err)
	}
}

func TestMessageRepo_ReadMessage(t *testing.T) {
	err := messageRepo.ReadMessage(uint(2), uint(1))

	if err != nil {
		t.Error(err)
	}
}
