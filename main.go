package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func handler(writer http.ResponseWriter, req *http.Request) {
	msg := req.FormValue("msg")
	if msg != "" {
		fmt.Println(req.FormValue("msg"))
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	task := &Task{Id: "1", Name: "GoでWeb APIを作る"}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(task); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(writer, buf.String())
}

func oldTodoHandler(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, "/todo", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/todo", handler)
	http.HandleFunc("/todo-old", oldTodoHandler)
	http.ListenAndServe(":8080", nil)
}
