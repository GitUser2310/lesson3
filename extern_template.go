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

var todos = []Todo{
	{"Выучить Go", false},
	{"Посетить лекцию по Go", false},
	{"...", false},
	{"Profit", false},
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

type Handler struct {
	http.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp1, err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	if r.Method == http.MethodPost {
		param := r.FormValue("id")
		index, _ := strconv.ParseInt(param, 10, 0)
		todos[index].Done = true
	}

	err1 := tmp1.Execute(w, todos)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
	}
}

func main() {
	handler := Handler{}
	http.Handle("/", handler)
	http.ListenAndServe(":8081", nil)
}
