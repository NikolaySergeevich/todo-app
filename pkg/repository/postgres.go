package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	USER_TABLE = "users"
	TODO_LIST_TABLE = "todo_lists"
	USERS_LIST_TABLE = "users_lists"
	TODO_ITEMS_TABLE = "todo_items"
	LISTS_ITEMS_TABLE = "lists_items"
)

type Config struct {
	Host 		string
	Port 		string
	UserName 	string
	Password 	string
	DBName 		string
	SSLMode 	string
}

func NewPostgresDB(ctf Config) (*sqlx.DB, error){
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
		ctf.Host, ctf.Port, ctf.UserName, ctf.DBName, ctf.Password, ctf.SSLMode))
	if err != nil{
		return nil, err
	}
	err = db.Ping()//Проверка подключения
	if err != nil{
		return nil, err
	}
	return db, nil
}