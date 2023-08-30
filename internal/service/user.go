package service

import (
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
	"github.com/linqcod/avito-internship-2023/internal/model"
)

type UserRepository interface {
	CreateUser(user dto.CreateUserDTO) (int64, error)
	GetUserById(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	AddUserToSegment(userId int64, slug string, ttl string) error
	DeleteUserFromSegment(userId int64, slug string) error
	GetUserActiveSegments(userId int64) ([]*model.UserSegment, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(user dto.CreateUserDTO) (int64, error) {
	return s.repo.CreateUser(user)
}

func (s UserService) GetUserById(id int64) (*model.User, error) {
	return s.repo.GetUserById(id)
}

func (s UserService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAllUsers()
}

// TODO: db transactions?
func (s UserService) ChangeUserSegments(userSegmDTO dto.ChangeUserSegmentsDTO, userId int64) error {
	for _, segm := range userSegmDTO.SegmentsToAdd {
		if err := s.repo.AddUserToSegment(userId, segm.Slug, segm.TTL); err != nil {
			return err
		}
	}
	for _, slug := range userSegmDTO.SegmentsToRemove {
		if err := s.repo.DeleteUserFromSegment(userId, slug); err != nil {
			return err
		}
	}

	return nil
}

func (s UserService) GetUserActiveSegments(userId int64) (*dto.ActiveUserSegmentsDTO, error) {
	segments, err := s.repo.GetUserActiveSegments(userId)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserSegmentToActiveUserSegments(userId, segments), nil
}
