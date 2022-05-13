package main

import (
	"fmt"
	"net/http"

	"github.com/pislas/activity_tracking/users"
)

// func handler(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
// }

/*func index(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(writer, "Hello World!")
}*/

/* type LoginForm struct {
	UserName string
	Password string
} */

func login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Fprintf(writer, "Debe ser un metodo Post")
		return
	}
	/* loginForm := LoginForm{
		UserName: request.FormValue("username"),
		Password: request.FormValue("password"),
	} */
	username := request.FormValue("username")
	password := request.FormValue("password")
	if password != "123" {
		fmt.Fprintf(writer, "Password invalido!")
		return
	}
	fmt.Fprintf(writer, "Bienvenido, %s!", username)
}

func main() {

	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", users.Register)
	http.ListenAndServe(":8081", nil)
}
