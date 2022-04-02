package main

// run --  go run main.go todo.go welcome.go

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func main() {
	selection := welcome()

	if selection == "1" {
		createNewTodo()
	} else if selection == "2" {
		editTodo()
	}
}

func createNewTodo() {
	inp, _ := getInput("enter a todo name :", reader)
	fmt.Println("new todo list created")

	td := newTodo(inp)
	promptInteractWithTodo(td)
}

func editTodo() {
	path, _ := os.Getwd()
	files, _ := ioutil.ReadDir(path + "/todo/")

	for i, file := range files {
		fmt.Printf("[%v]...%v\n", i, file.Name())
	}

	inp, _ := getInput("please input your todo index:", reader)
	index, err := strconv.ParseInt(inp, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(path + "/todo/" + files[index].Name())
	if err != nil {
		log.Fatal(err)
	}

	todo := string(content)
	todo = strings.Split(todo, "\n\n")[1]
	todoName := strings.Split(todo, "\n")[0]
	todoValue := strings.Split(todo, "\n")[1:]

	edittd := newTodo(todoName)

	for _, v := range todoValue[:len(todoValue)-1] {
		curVal := strings.Split(v, "-")
		// fmt.Println(curVal[2])
		var done bool

		if curVal[2] == "done" {
			done = true
		} else {
			done = false
		}

		newList := TodoList{
			value: curVal[1],
			done:  done,
		}

		edittd.todoList = append(edittd.todoList, newList)
	}

	promptInteractWithTodo(edittd)
}

func promptInteractWithTodo(todo Todo) {
	inp, _ := getInput("(a - add to list, r - remove from list, d - set task as done, s - save todo) :", reader)

	switch inp {
	case "a":
		tdVal, _ := getInput("type in your todo :", reader)
		todo.addToList(tdVal)
		promptInteractWithTodo(todo)

	case "r":
		tdNumber, _ := getInput("enter the todo index:", reader)
		todo.removeFromList(tdNumber)
		promptInteractWithTodo(todo)

	case "d":
		tdNumber, _ := getInput("enter the todo index:", reader)
		mark, _ := getInput("enter 't' or 'f' to mark as done or not done:", reader)
		todo.markTodo(tdNumber, mark)
		promptInteractWithTodo(todo)

	case "s":
		todo.saveTodo()

	default:
		promptInteractWithTodo(todo)
	}
}
