package repository

import (
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var db = datasource.InstanceGormMaster()
var collectTopicRepo = NewCollectTopicRepo(datasource.InstanceGormMaster())

func TestCollectTopicRepo_Collect(t *testing.T) {
	var topic = new(entitys.Topic)
	var user = new(entitys.User)
	db.First(topic)
	db.First(user)
}
