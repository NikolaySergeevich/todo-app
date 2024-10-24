package service

import (
	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/NikolaySergeevich/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (todoListSer *TodoListService) Create(userId int, todolist todoapp.TodoList) (int, error) {
	return todoListSer.repo.Create(userId, todolist)
}