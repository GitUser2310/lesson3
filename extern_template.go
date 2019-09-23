package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

func main() {
	tmp1, err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}
	todos := []Todo{
		{"Выучить Go", false},
		{"Посетить лекцию по Go", false},
		{"...", false},
		{"Profit", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			param := r.FormValue("id")
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true
		}

		err := tmp1.Execute(w, todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8081", nil)
}
