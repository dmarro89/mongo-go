package service

import (
	"context"
	"mongo-go/models"
)

type ServiceMock struct {
	Error error
	User  *models.User
	Users []models.User
}

func NewMock() Service {
	return &ServiceMock{}
}

func NewErrMock(err error) Service {
	return &ServiceMock{Error: err}
}

func (s *ServiceMock) Ping(ctx context.Context) error {
	return s.Error
}

func (s *ServiceMock) Insert(ctx context.Context, user models.User) error {
	return s.Error
}

func (s *ServiceMock) Update(ctx context.Context, user models.User) error {
	return s.Error
}

func (s *ServiceMock) GetByID(ctx context.Context, userId string) (*models.User, error) {
	return s.User, s.Error
}

func (s *ServiceMock) Get(ctx context.Context) ([]models.User, error) {
	return s.Users, s.Error
}

func (s *ServiceMock) DeleteByID(ctx context.Context, userId string) error {
	return s.Error
}
