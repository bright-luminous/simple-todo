package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type JSON map[string]interface{}

func writeJSON(rw http.ResponseWriter, status int, data JSON) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}

func echoRequest(rw http.ResponseWriter, r *http.Request) {
	type response struct {
		Method string      `json:"method"`
		Header interface{} `json:"header"`
		Body   interface{} `json:"body"`
		Query  interface{} `json:"query"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	requestBody := map[string]interface{}{}
	if len(bytes) != 0 {
		if err := json.Unmarshal(bytes, &requestBody); err != nil {
			panic(err.Error())
		}
	}

	writeJSON(rw, http.StatusOK, JSON{
		"data": &response{
			Method: r.Method,
			Header: r.Header,
			Body:   requestBody,
			Query:  r.URL.Query(),
		},
	})
}

func getTodos(rw http.ResponseWriter, r *http.Request) {
	todos := []*Todo{}

	for _, todo := range todoRepository {
		todos = append(todos, todo)
	}

	writeJSON(rw, http.StatusOK, JSON{
		"data": todos,
	})
}

func getTodoByID(rw http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	for id, todo := range todoRepository {
		if id == todoID {
			writeJSON(rw, http.StatusOK, JSON{
				"data": todo,
			})
		}
	}

}

type createTodoRequestBody struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}

func createTodo(rw http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		writeJSON(rw, http.StatusInternalServerError, JSON{
			"data":  nil,
			"error": "internal server error",
		})
		return
	}

	requestBody := &createTodoRequestBody{}
	if err := json.Unmarshal(bytes, requestBody); err != nil {
		log.Println(err.Error())
		writeJSON(rw, http.StatusInternalServerError, JSON{
			"data":  nil,
			"error": "internal server error",
		})
		return
	}

	todo := NewTodo(requestBody.Task, requestBody.Description)
	todoRepository[todo.ID.String()] = todo

	writeJSON(rw, http.StatusOK, JSON{
		"data": todo,
	})
}

func deleteTodo(rw http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	for id := range todoRepository {
		if id == todoID {
			delete(todoRepository, todoID)
			rw.WriteHeader(http.StatusNoContent)
			return
		}
	}

	rw.WriteHeader(http.StatusNotFound)
}

type updateTodoRequestBody struct {
	Task        *string `json:"task"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func updateTodo(rw http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		writeJSON(rw, http.StatusInternalServerError, JSON{
			"data":  nil,
			"error": "internal server error",
		})
		return
	}
	requestBody := &updateTodoRequestBody{}
	if err := json.Unmarshal(bytes, requestBody); err != nil {
		panic(err.Error())
	}

	var foundTodo *Todo
	for id, todo := range todoRepository {
		if id == todoID {
			todo.Update(requestBody.Task, requestBody.Description, requestBody.Done)
			foundTodo = todo
		}
	}

	writeJSON(rw, http.StatusOK, JSON{
		"data": foundTodo,
	})

}
