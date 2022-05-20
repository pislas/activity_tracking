package users

import (
	"fmt"
)

var UsersDb *Users

type User struct {
	Id       uint64
	Username string
	Password string
}

type Users struct {
	users []User
}

func init() {
	UsersDb = &Users{}
}

func (s *Users) Insert(u User) User {
	u.Id = uint64(len(s.users))
	s.users = append(s.users, u)
	return u
}

func (s *Users) FindUser(username string) (*User, error) {
	for i, u := range s.users {
		fmt.Printf("Indice %d => %v", i, u)
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("usuario %s invalido", username)
}
