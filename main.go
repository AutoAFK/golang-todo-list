package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AutoAFK/golang-todo-list/manager"
)

func main() {
	fmt.Println("Todo list:")
	fmt.Println("==========")
	fmt.Println("1. Create new task")
	fmt.Println("2. Edit task")
  fmt.Println("3. Delete task")
	fmt.Println("4. View task list")
	fmt.Println("0. Exit")

	manager := manager.NewManager()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter an option: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong in input :(")
			return
		}
		text = strings.TrimSpace(text)
		switch text {
		case "1":
			manager.CreateNewTask()
		case "2":
      err := manager.EditTask()
      if err != nil {
        panic(err)
      }
    case "3":
      err := manager.DeleteTask()
      if err != nil {
        panic(err)
      }
		case "4":
			manager.ViewTasks()
		case "0", "exit":
			return
		default:
			fmt.Println("Please enter a valid option.")
		}
	}
}
