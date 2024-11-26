package repository

import (
	"fmt"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{ db: db}
}

func(itemPostgres *TodoItemPostgres) Create(listId int, item todoapp.TodoItem) (int, error) {
	tx, err := itemPostgres.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	// это запрос для заполнения таблицы задач
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", TODO_ITEMS_TABLE)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// это запрос для заполнения таблицв для связки списка и конкртной задачи
	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2) RETURNING id", LISTS_ITEMS_TABLE)

	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, nil
}