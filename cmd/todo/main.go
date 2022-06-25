package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/gyu-young-park/what-to-do/api"
)

const (
	TODO_FILE = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "make a todo as completed")
	del := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "lists all todo in list")

	flag.Parse()

	todoList := &todo.TodoList{}

	if err := todoList.Load(TODO_FILE); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := GetInput(os.Stdin, flag.Args()...)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}
		todoList.Add(task)
		err = todoList.Store(TODO_FILE)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}
	case *complete > 0:
		err := todoList.Complete(*complete)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}
		err = todoList.Store(TODO_FILE)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}
	case *del > 0:
		err := todoList.Delete(*del)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}

		err = todoList.Store(TODO_FILE)
		if err != nil {
			ErrorReportAndSystemExit(err)
		}
	case *list:
		todoList.List()
	}
}
