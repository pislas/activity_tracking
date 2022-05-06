package main

import (
	"net/http"
)

// func handler(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
// }

/*func index(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(writer, "Hello World!")
}*/

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8081", nil)
}
