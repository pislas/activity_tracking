package users

import (
	"fmt"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Fprintf(writer, "Debe ser un metodo Post")
		return
	}
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "Datos invalidos para registro!")
		return
	}
	user := User{
		Id:       uint64(len(UsersDb.users)),
		Username: request.PostFormValue("username"),
		Password: request.PostFormValue("password"),
	}
	user = UsersDb.Insert(user)
	fmt.Fprintf(writer, "Usuario Registrado nro: %d, Bienvenido %s!", len(UsersDb.users), user.Username)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Fprintf(writer, "Debe ser un metodo Post")
		return
	}
	username := request.FormValue("username")
	u, error := UsersDb.FindUser(username)
	if error != nil {
		//fmt.Printf("error: %v\n", error)
		fmt.Fprintf(writer, "Usuario invalido")
		return
	}
	password := request.FormValue("password")
	if password != u.Password {
		fmt.Fprintf(writer, "Password invalido!")
		return
	}
	fmt.Fprintf(writer, "Bienvenido, %s!", username)
}
