package main

import (
	"log"

	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/NikolaySergeevich/todo-app/pkg/handler"
	"github.com/NikolaySergeevich/todo-app/pkg/repository"
	"github.com/NikolaySergeevich/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err :=initConfig(); err != nil{
		log.Fatalf("error initalizing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs") //Имя дерриктории
	viper.SetConfigName("config") // Имя файла
	return viper.ReadInConfig()   // Возвращает функцию, которая считывает значение конфига и записывает их во внутренний объект viper
}
