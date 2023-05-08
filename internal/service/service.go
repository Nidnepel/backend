package service

import "github.com/Nidnepel/backend/internal/repository"

type Authorization interface {
}

type Project interface {
}

type Task interface {
}

type Manager interface {
}

type Worker interface {
}

type Service struct {
	Authorization
	Project
	Task
	Manager
	Worker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
