package repository

import "github.com/linqcod/avito-internship-2023/internal/domain/model"

type UserRepository interface {
	CreateUser(user model.CreateUserDTO) (int64, error)
	GetUserById(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	AddUserToSegment(userId int64, slug string) error
	DeleteUserFromSegment(userId int64, slug string) error
	GetUserActiveSegments(userId int64) (*model.ActiveUserSegmentsDTO, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(user model.CreateUserDTO) (int64, error) {
	return s.repo.CreateUser(user)
}

func (s UserService) GetUserById(id int64) (*model.User, error) {
	return s.repo.GetUserById(id)
}

func (s UserService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAllUsers()
}

func (s UserService) AddUserToSegment(userId int64, slug string) error {
	return s.repo.AddUserToSegment(userId, slug)
}

func (s UserService) DeleteUserFromSegment(userId int64, slug string) error {
	return s.repo.DeleteUserFromSegment(userId, slug)
}

func (s UserService) GetUserActiveSegments(userId int64) (*model.ActiveUserSegmentsDTO, error) {
	return s.repo.GetUserActiveSegments(userId)
}
