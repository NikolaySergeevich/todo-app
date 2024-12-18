package repository

import (
	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(user todoapp.User) (int, error)
	GetUser(userName, password string) (todoapp.User, error)
}
type TodoList interface{
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, Id int) (todoapp.TodoList, error)
	Delete(userId, Id int) error
	Update(userId, idList int, updatePayload todoapp.UpdateListPayload) error
}
type TodoItem interface{
	Create(listId int, item todoapp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoapp.TodoItem, error)
}

type Repository struct{
	Authorization
	TodoList
	TodoItem
}
func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}