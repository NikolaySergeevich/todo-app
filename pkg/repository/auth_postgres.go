package repository

import (
	"fmt"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (authPostgres *AuthPostgres) CreateUser(user todoapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", USER_TABLE)
	row := authPostgres.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (authPostgres *AuthPostgres) GetUser(userName, password string) (todoapp.User, error) {
	var user todoapp.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", USER_TABLE)
	err := authPostgres.db.Get(&user, query, userName, password)
	return user, err
}