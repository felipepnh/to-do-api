package repository

import "to-do-api/internal/domain"

type Todo interface {
	Insert(todo domain.Todo) (domain.Todo, error)
	GetAll() ([]domain.Todo, error)
	GetById(id string) (domain.Todo, error)
	DeleteById(id string) error
	UpdateById(id string, todo domain.Todo) (domain.Todo, error)
}
