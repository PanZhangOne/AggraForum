package repository

import (
	"forum/datasource"
	"forum/entitys"
	"testing"
)

var labelsRepo = NewLabelsRepo(datasource.InstanceGormMaster())

func TestLabelsRepo_Create(t *testing.T) {
	var label = new(entitys.Label)
	label.LabelName = "test"

	err := labelsRepo.Create(label)
	if err != nil {
		t.Error(err)
	}

	labelsRepo.db.Unscoped().Delete(label)
}

func TestLabelsRepo_FindByID(t *testing.T) {
	var label = new(entitys.Label)
	label.LabelName = "test"

	err := labelsRepo.Create(label)
	if err != nil {
		t.Error(err)
	}

	_l, err := labelsRepo.FindByID(label.ID)
	if err != nil {
		labelsRepo.db.Unscoped().Delete(label)
		t.Error(err)
	}
	if _l.ID != label.ID || _l.LabelName != label.LabelName {
		labelsRepo.db.Unscoped().Delete(label)
		t.Error("查找失败")
	}

	labelsRepo.db.Unscoped().Delete(label)
}

func TestLabelsRepo_FindByIDs(t *testing.T) {
	var (
		label1 = new(entitys.Label)
		label2 = new(entitys.Label)
		label3 = new(entitys.Label)
	)

	label1.LabelName = "test1"
	label2.LabelName = "test2"
	label3.LabelName = "test3"

	err := labelsRepo.Create(label1)
	if err != nil {
		t.Error(err)
	}
	err = labelsRepo.Create(label2)
	if err != nil {
		t.Error(err)
	}
	err = labelsRepo.Create(label3)
	if err != nil {
		t.Error(err)
	}

	ids := make([]uint, 3)
	ids = append(ids, label1.ID, label2.ID, label3.ID)
	labels, err := labelsRepo.FindByIDs(ids)
	if err != nil {
		t.Error(err)
	}
	if len(labels) != 3 {
		t.Error("查找失败")
	}
	labelsRepo.db.Unscoped().Delete(labels)
}

func TestLabelsRepo_Update(t *testing.T) {
	label := entitys.Label{
		LabelName: "test",
	}
	err := labelsRepo.Create(&label)
	if err != nil {
		t.Error(err)
	}
	label.LabelName = "change_test"
	err = labelsRepo.Update(&label)
	if err != nil {
		t.Error("更新失败:", err)
		labelsRepo.db.Unscoped().Delete(&label)
		return
	}

	_l, err := labelsRepo.FindByID(label.ID)
	if err != nil {
		labelsRepo.db.Unscoped().Delete(&label)
		t.Error("查找失败:", err)
	}
	if _l.ID != label.ID || _l.LabelName != label.LabelName {
		labelsRepo.db.Unscoped().Delete(&label)
		t.Error("更新失败")
	}
	labelsRepo.db.Unscoped().Delete(&label)
}
