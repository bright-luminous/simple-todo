package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var todoRepository = map[string]*Todo{}

func init() {
	starterTodos := []*Todo{
		NewTodo("Learn how to code Go", "Go is easy!"),
		NewTodo("Buy coffee", "I'm so sleepy!"),
		NewTodo("Clean the room", "Is this really a human room?"),
	}

	for _, todo := range starterTodos {
		todoRepository[todo.ID.String()] = todo
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Handle("/", http.HandlerFunc(echoRequest))
	})

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getTodos)
		r.Get("/{id}", getTodoByID)
		r.Post("/", createTodo)
		r.Delete("/{id}", deleteTodo)
		r.Put("/{id}", updateTodo)
	})

	fmt.Println("server is listening at port 3000!")
	http.ListenAndServe(":3000", r)
}
