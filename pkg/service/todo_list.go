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

func (todoListSer *TodoListService) GetAll(userId int) ([]todoapp.TodoList, error) {
	return todoListSer.repo.GetAll(userId)
}

func (todoListSer *TodoListService) GetById(userId, Id int) (todoapp.TodoList, error) {
	return todoListSer.repo.GetById(userId, Id)
}

func (todoListSer *TodoListService) Delete(userId, Id int) ( error) {
	return todoListSer.repo.Delete(userId, Id)
}

func (todoListSer *TodoListService) Update(userId, Id int, updatePayload todoapp.UpdateListPayload) ( error) {
	if err := updatePayload.Validate(); err != nil {
		return err
	}
	return todoListSer.repo.Update(userId, Id, updatePayload)
}