package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTodo(t *testing.T) {
	t.Run("shold create a todo", func(t *testing.T) {
		expectedTitle := "Title"
		expectedDescription := "Description"

		todo, err := NewTodo(expectedTitle, expectedDescription)

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, expectedTitle, todo.Title)
		assert.Equal(t, expectedDescription, todo.Description)
	})
	t.Run("shold create a empty todo", func(t *testing.T) {
		expectedTitle := ""
		expectedDescription := ""

		todo, err := NewTodo(expectedTitle, expectedDescription)
		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, ErrTodoIsInvalid.Error(), err.Error())

	})
}
