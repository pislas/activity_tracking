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
		Username: request.PostFormValue("name"),
		Password: request.PostFormValue("password"),
	}
	user = UsersDb.Insert(user)
	fmt.Fprintf(writer, "Usuario Registrado nro: %d, Bienvenido %s!", len(UsersDb.users), user.Username)
}
