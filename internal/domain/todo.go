package domain

import (
	"errors"
)

var ErrTodoIsInvalid = errors.New("Todo is invalid")

type Todo struct {
	ID          string
	Title       string
	Description string
}

func NewTodo(title, description string) (Todo, error) {
	if title == "" || description == "" {
		return Todo{}, ErrTodoIsInvalid
	}
	return Todo{
		Title:       title,
		Description: description,
	}, nil
}
