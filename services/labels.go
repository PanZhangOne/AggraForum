package services

import (
	"forum/datasource"
	"forum/entitys"
	"forum/repository"
)

type LabelsService interface {
	Create(label *entitys.Labels) error

	FindAll(limit, offset uint) ([]entitys.Labels, error)
	FindByID(id uint) (*entitys.Labels, error)
	FindByIDs(ids []uint) ([]entitys.Labels, error)

	FindAllLabel() []entitys.Labels
	FindHotLabels() []entitys.Labels

	PostTopicHandle(label *entitys.Labels)
}

type labelService struct {
	repo *repository.LabelsRepo
}

func (s *labelService) Create(label *entitys.Labels) error {
	return s.repo.Create(label)
}

func (s *labelService) FindAll(limit, offset uint) ([]entitys.Labels, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *labelService) FindAllLabel() []entitys.Labels {
	labels, _ := s.repo.FindAllLabels()
	return labels
}

func (s *labelService) FindByID(id uint) (*entitys.Labels, error) {
	return s.repo.FindByID(id)
}

func (s *labelService) FindByIDs(ids []uint) ([]entitys.Labels, error) {
	return s.repo.FindByIDs(ids)
}

func (s *labelService) FindHotLabels() []entitys.Labels {
	labels, _ := s.repo.FindHotLabels()
	return labels
}

// PostTopicHandle
func (s *labelService) PostTopicHandle(label *entitys.Labels) {
	label.TopicsCount += 1
	_ = s.repo.Update(label)
}

func NewLabelService() LabelsService {
	return &labelService{
		repo: repository.NewLabelsRepo(datasource.InstanceGormMaster()),
	}
}
