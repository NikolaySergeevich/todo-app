package todoapp

import (
	"encoding/json"
	"errors"
)

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

func (todo *TodoList) MarshalJSON() ([]byte, error) {
	type Alias TodoList
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	})
}

func (todo *TodoList) UnmarshalJSON(b []byte) error {
	type Alias TodoList
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	}
	return json.Unmarshal(b, &aux)
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (todo *TodoItem) MarshalJSON() ([]byte, error) {
	type Alias TodoItem
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	})
}

func (todo *TodoItem) UnmarshalJSON(b []byte) error {
	type Alias TodoItem
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	}
	return json.Unmarshal(b, &aux)
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListPayload struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (todo *UpdateListPayload) MarshalJSON() ([]byte, error) {
	type Alias UpdateListPayload
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	})
}

func (todo *UpdateListPayload) UnmarshalJSON(b []byte) error {
	type Alias UpdateListPayload
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	}
	return json.Unmarshal(b, &aux)
}

func (todo *UpdateListPayload) Validate() error {
	if todo.Title == nil && todo.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
