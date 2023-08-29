package repository

import "github.com/linqcod/avito-internship-2023/internal/domain/model"

type SegmentRepository interface {
	CreateSegment(segment model.CreateSegmentDTO) (string, error)
	DeleteSegment(slug string) error
}

type SegmentService struct {
	repo SegmentRepository
}

func NewSegmentService(repo SegmentRepository) *SegmentService {
	return &SegmentService{
		repo: repo,
	}
}

func (s SegmentService) CreateSegment(segment model.CreateSegmentDTO) (string, error) {
	return s.repo.CreateSegment(segment)
}

func (s SegmentService) DeleteSegment(slug string) error {
	return s.repo.DeleteSegment(slug)
}
