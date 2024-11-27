package service

import (
	todoapp "github.com/NikolaySergeevich/todo-app"
	"github.com/NikolaySergeevich/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	lsitRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, repoList repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, lsitRepo: repoList}
}

func(item *TodoItemService) Create(userId, listID int, itemPayload todoapp.TodoItem) (int, error) {
	_, err := item.lsitRepo.GetById(userId, listID)
	if err != nil {
		// список с таким id не найден или не соответствует с пользователем
		return 0, err
	}

	return item.repo.Create(listID, itemPayload)
}

func(item *TodoItemService) GetAll(userId, listId int) ([]todoapp.TodoItem, error) {
	return item.repo.GetAll(userId, listId)
}