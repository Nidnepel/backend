package repository

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

type Repository struct {
	Authorization
	Project
	Task
	Manager
	Worker
}

func NewRepository() *Repository {
	return &Repository{}
}
