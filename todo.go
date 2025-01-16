package main

import (
	"errors"
	"time"
	"fmt"
)

type Todo struct {
	Tittle string
	Compleated bool
	CreatedAt time.Time
	CompleatedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo {
		Tittle: title,
		Compleated: false,
		CompleatedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}