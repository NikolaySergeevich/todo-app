package handler

import (
	"github.com/NikolaySergeevich/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	services *service.Service
}

func NewHandler(servise *service.Service) *Handler{
	return &Handler{services: servise}
}

// Инициализирует все andpoints
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()// Инициализация роутера

	// Объявление методов, сгруппировав их по маршрутам:

	// В группе /auth описывется два andpoints для регистрации и для авторизации
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)//для регистрации
		auth.POST("/sign-in", h.signIn)//для авторизации
	}
	// Группа /api используется для работы andpoints со списками и их задачами
	api := router.Group("/api")
	{
		// Внутри группы api создаётся группа lists для работы со списками:
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createlist)
			lists.GET("/", h.getAlllist)
			lists.GET("/:id", h.getlistById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// Создаётся группа задач списка
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}