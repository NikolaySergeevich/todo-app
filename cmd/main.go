package main

import (
	"os"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/NikolaySergeevich/todo-app/pkg/handler"
	"github.com/NikolaySergeevich/todo-app/pkg/repository"
	"github.com/NikolaySergeevich/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err :=initConfig(); err != nil{
		logrus.Fatalf("error initalizing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil{//выгружаем переменные окружения из файла
		logrus.Fatalf("error loading env varibales: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"), //смотри какой порт указывал при запуске контейнера
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),//читаем переменные окр. из системы
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil{
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs") //Имя дерриктории
	viper.SetConfigName("config") // Имя файла
	return viper.ReadInConfig()   // Возвращает функцию, которая считывает значение конфига и записывает их во внутренний объект viper
}
