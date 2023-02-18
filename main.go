package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/todo", Handler)
	http.HandleFunc("/todo-old", OldTodoHandler)
	http.ListenAndServe(":8080", nil)
}
