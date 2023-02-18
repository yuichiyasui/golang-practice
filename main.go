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

func handler(w http.ResponseWriter, r *http.Request) {
	task := &Task{Id: "1", Name: "GoでWeb APIを作る"}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(task); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
