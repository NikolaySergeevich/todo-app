package todoapp

import "encoding/json"

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (todo *TodoList) MarshalJSON() ([]byte, error)  {
	type Alias TodoList
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(todo),
	})
}

func (todo *TodoList) UnmarshalJSON(b []byte) error {
	type Alias TodoList
	aux := &struct{
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
	return json.Marshal(&struct{
		*Alias
	}{
		Alias: (*Alias)(todo),
	})
}

func (todo *TodoItem) UnmarshalJSON(b []byte) error {
	type Alias TodoItem
	aux := &struct{
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
