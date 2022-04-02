package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func welcome() string {
	var in string
	test := true

	for test {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		fmt.Println("=====================")
		fmt.Println("your to do app")
		fmt.Println("=====================")
		fmt.Println("")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("[1] create a new to-do list")
		fmt.Println("[2] edit an existing to-do list")
		input, _ := getInput("select one to continue:", reader)

		if input == "1" || input == "2" {
			test = false
			in = input
		} else {
			test = true
		}
	}

	return in
}
