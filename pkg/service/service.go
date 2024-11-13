package service

import (
	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/NikolaySergeevich/todo-app/pkg/repository"
)

type Authorization interface{
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface{
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, Id int) (todoapp.TodoList, error)
	Delete(userId, id int) error
	Update(userId, idList int, updatePayload todoapp.UpdateListPayload) error
}
type TodoItem interface{

}

type Service struct{
	Authorization
	TodoList
	TodoItem
}
func NewService(repo repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList: NewTodoListService(repo.TodoList),
	}
}