package users

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
