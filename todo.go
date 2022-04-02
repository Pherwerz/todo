package main

import (
	"fmt"
	"os"
	"strconv"
)

type TodoList struct {
	value string
	done  bool
}

type Todo struct {
	name     string
	todoList []TodoList
}

func newTodo(n string) Todo {
	todo := Todo{
		name:     n,
		todoList: []TodoList{},
	}

	return todo
}

func format(td *Todo) string {
	fmt.Println("")
	fs := "current todo list \n\n"
	fs += fmt.Sprintf("%v\n", td.name)

	for i, v := range td.todoList {
		var done string

		if v.done {
			done = "done"
		} else {
			done = "not done"
		}

		fs += fmt.Sprintf("[%v]-%-25v -%v\n", i, v.value, done)
	}

	fmt.Println(fs)

	return fs
}

func (todo *Todo) addToList(todoValue string) {
	newList := TodoList{
		value: todoValue,
		done:  false,
	}

	todo.todoList = append(todo.todoList, newList)
	format(todo)
}

func (todo *Todo) removeFromList(id string) {
	index, _ := strconv.ParseInt(id, 10, 32)

	todo.todoList = append(todo.todoList[:index], todo.todoList[index+1:]...)
	format(todo)
}

func (todo *Todo) markTodo(id string, mark string) {
	index, _ := strconv.ParseInt(id, 10, 32)

	if mark == "t" {
		todo.todoList[index].done = true
	} else {
		todo.todoList[index].done = false
	}
	format(todo)
}

func (todo *Todo) saveTodo() {
	data := []byte(format(todo))
	os.WriteFile("todo/"+todo.name+".txt", data, 0644)
}
