package repository

import (
	"fmt"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{ db: db}
}

func (dataB *TodoListPostgres) Create(userId int, list todoapp.TodoList) (int, error) {
	tx, err := dataB.db.Begin()
	if err != nil {
		return 0, err
	}

	// подготовка и исполнение запроса на запись данных с получением в ответ id записи
	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", TODO_LIST_TABLE)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	// подготовка и исполнение запроса для добавление записи в таблицу без возвращения каких-то данынх
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", USERS_LIST_TABLE)
	_, err = tx.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}