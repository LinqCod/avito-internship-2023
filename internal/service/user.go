package service

import (
	model2 "github.com/linqcod/avito-internship-2023/internal/model"
)

type UserRepository interface {
	CreateUser(user model2.CreateUserDTO) (int64, error)
	GetUserById(id int64) (*model2.User, error)
	GetAllUsers() ([]*model2.User, error)
	AddUserToSegment(userId int64, slug string) error
	DeleteUserFromSegment(userId int64, slug string) error
	GetUserSegments(userId int64) ([]*model2.UserSegment, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(user model2.CreateUserDTO) (int64, error) {
	return s.repo.CreateUser(user)
}

func (s UserService) GetUserById(id int64) (*model2.User, error) {
	return s.repo.GetUserById(id)
}

func (s UserService) GetAllUsers() ([]*model2.User, error) {
	return s.repo.GetAllUsers()
}

// TODO: create errors, db transactions?, !!ADD LOGS!!
func (s UserService) ChangeUserSegments(dto model2.ChangeUserSegmentsDTO, userId int64) error {
	for _, slug := range dto.SegmentsToAdd {
		if err := s.repo.AddUserToSegment(userId, slug); err != nil {
			return err
		}
	}
	for _, slug := range dto.SegmentsToRemove {
		if err := s.repo.DeleteUserFromSegment(userId, slug); err != nil {
			return err
		}
	}

	return nil
}

func (s UserService) GetUserActiveSegments(userId int64) (*model2.ActiveUserSegmentsDTO, error) {
	segments, err := s.repo.GetUserSegments(userId)
	if err != nil {
		return nil, err
	}

	return model2.ConvertUserSegmentToActiveUserSegments(userId, segments), nil
}
