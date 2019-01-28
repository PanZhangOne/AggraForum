package repository

import (
	"forum/entitys"
	"github.com/jinzhu/gorm"
)

type LabelsRepo struct {
	db *gorm.DB
}

func NewLabelsRepo(db *gorm.DB) *LabelsRepo {
	return &LabelsRepo{db: db}
}

func (r *LabelsRepo) Create(label *entitys.Label) error {
	return r.db.Create(label).Error
}

func (r *LabelsRepo) FindByID(id uint) (*entitys.Label, error) {
	var label = new(entitys.Label)
	label.ID = id
	err := r.db.First(label).Error
	return label, err
}

func (r *LabelsRepo) FindByIDs(ids []uint) ([]entitys.Label, error) {
	var labels = make([]entitys.Label, 0)
	err := r.db.Where("id in (?)", ids).Find(&labels).Error
	return labels, err
}

func (r *LabelsRepo) FindAll(limit, offset uint) ([]entitys.Label, error) {
	var labels = make([]entitys.Label, 0)
	err := r.db.Find(&labels).Limit(limit).Offset(offset).Error
	return labels, err
}

func (r *LabelsRepo) FindAllLabels() ([]entitys.Label, error) {
	var labels = make([]entitys.Label, 0)
	err := r.db.Find(&labels).Error
	return labels, err
}

func (r *LabelsRepo) FindHotLabels() ([]entitys.Label, error) {
	var labels = make([]entitys.Label, 0)
	err := r.db.Order("topics_count desc").Limit(10).Find(&labels).Error
	return labels, err
}

func (r *LabelsRepo) Update(label *entitys.Label) error {
	return r.db.Save(label).Error
}
