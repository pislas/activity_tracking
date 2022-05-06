package main

import (
	"fmt"
	"net/http"
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
	fmt.Fprintf(writer, "Bienvenido, %s!", username)
}

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8081", nil)
}
