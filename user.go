package todoapp

import "encoding/json"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Marshal() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(user),
	})
}

func (user *User) Unmarshal(b []byte) error {
	type Alias User
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(user),
	}
	return json.Unmarshal(b, &aux)
}
