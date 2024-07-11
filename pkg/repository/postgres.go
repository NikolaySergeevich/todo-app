package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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