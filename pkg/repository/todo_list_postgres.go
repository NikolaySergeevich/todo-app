package repository

import (
	"fmt"
	"strings"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (dataB *TodoListPostgres) GetAll(userId int) ([]todoapp.TodoList, error) {
	var lists []todoapp.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", 
		TODO_LIST_TABLE, USERS_LIST_TABLE)

	err := dataB.db.Select(&lists, query, userId)
	return lists, err
}

func (dataB *TodoListPostgres) GetById(userId, Id int) (todoapp.TodoList, error) {
	var list todoapp.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul 
						on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, TODO_LIST_TABLE, USERS_LIST_TABLE)

	err := dataB.db.Get(&list, query, userId, Id)
	return list, err
}

func (dataB *TodoListPostgres) Delete(userId, Id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", TODO_LIST_TABLE, USERS_LIST_TABLE)

	_, err := dataB.db.Exec(query, userId, Id)
	return err
}

func (dataB *TodoListPostgres) Update(userId, Id int, updatePayload todoapp.UpdateListPayload) error {
	setValues := make([]string, 0) // в этот слайс будем складывать кусочек строки для запроса, который указвает какойму полю какое ххначение присваивать.
	args := make([]interface{}, 0) // в этом слайсе будут лежать наши новые значения, если они есть в payload.
	argsId := 1 // это будет задавть id для кусочка строки для обновления, которая будет лежать в первом слайсе(setValues)

	if updatePayload.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *updatePayload.Title)
		argsId++
	}
	if updatePayload.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *updatePayload.Description)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
			TODO_LIST_TABLE, setQuery, USERS_LIST_TABLE, argsId, argsId+1)
	
	args = append(args, userId, Id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args...)

	_, err := dataB.db.Exec(query, args...)

	return err
}