package users

import (
	"fmt"
)

var UsersDb *Users

type Users struct {
	users []User
}

type User struct {
	Id       uint64
	Username string
	Password string
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
		//fmt.Printf("Indice %d => %v", i, u)
		if u.Username == username {
			return &s.users[i], nil
		}
	}
	return nil, fmt.Errorf("usuario %s invalido", username)
}

func (s *Users) GetUser(id uint64) (*User, error) {
	for i, u := range s.users {
		if u.Id == id {
			return &s.users[i], nil
		}
	}
	return nil, fmt.Errorf("id de usuario %v invalido", id)
}

/* func (s *Users) UpdateUser(id uint64, username string, password string) (User, error) {
	_, u := UsersDb.GetUser(id)
} */
