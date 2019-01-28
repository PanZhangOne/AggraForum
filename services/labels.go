package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
)

type LabelsService interface {
	Create(label *entitys.Label) error

	FindAll(limit, offset uint) ([]entitys.Label, error)
	FindByID(id uint) (*entitys.Label, error)
	FindByIDs(ids []uint) ([]entitys.Label, error)

	FindAllLabel() []entitys.Label
	FindHotLabels() []entitys.Label

	PostTopicHandle(label *entitys.Label)
}

type labelService struct {
	repo *repository.LabelsRepo
}

func (s *labelService) Create(label *entitys.Label) error {
	return s.repo.Create(label)
}

func (s *labelService) FindAll(limit, offset uint) ([]entitys.Label, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *labelService) FindAllLabel() []entitys.Label {
	labels, _ := s.repo.FindAllLabels()
	return labels
}

func (s *labelService) FindByID(id uint) (*entitys.Label, error) {
	return s.repo.FindByID(id)
}

func (s *labelService) FindByIDs(ids []uint) ([]entitys.Label, error) {
	return s.repo.FindByIDs(ids)
}

func (s *labelService) FindHotLabels() []entitys.Label {
	labels, _ := s.repo.FindHotLabels()
	return labels
}

// PostTopicHandle
func (s *labelService) PostTopicHandle(label *entitys.Label) {
	label.TopicsCount += 1
	_ = s.repo.Update(label)
}

func NewLabelService() LabelsService {
	return &labelService{
		repo: repository.NewLabelsRepo(datasource.InstanceGormMaster()),
	}
}
