package main

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Task        string    `json:"task"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Done        bool      `json:"done"`
}

func (t *Todo) Update(task *string, description *string, done *bool) {
	if task != nil {
		t.Task = *task
	}

	if description != nil {
		t.Description = *description
	}

	if done != nil {
		t.Done = *done
	}
}

func NewTodo(givenTask string, givenDescription string) *Todo {
	return &Todo{
		ID:          uuid.New(),
		Task:        givenTask,
		Description: givenDescription,
		CreatedAt:   time.Now(),
		Done:        false,
	}
}
