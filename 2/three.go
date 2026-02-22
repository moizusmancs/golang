package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TodoApp struct {
	tasks []string
	scanner *bufio.Scanner
}

func NewTodoApp() *TodoApp {
	return &TodoApp{
		tasks: []string{},
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (app *TodoApp) Run(){
	for {
		app.printMenu()
		choice := app.readInput()

		if choice == "" {
			fmt.Println("Choice cannot be empty. Please choose a valid option.")
			continue
		}

		switch choice {
		case "1":
			app.addTask()
		case "2":
			app.viewTasks()
		case "3":
			app.deleteTask()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, choose again.")
		}
	}
}

func (app *TodoApp) printMenu(){
	fmt.Println("\n--- Todo App ---")
	fmt.Println("1. Add a task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Tasks")
	fmt.Println("4. Exit")
	fmt.Print("Choose an option: ")
}

func (app *TodoApp) readInput() string {
	app.scanner.Scan()
	return strings.TrimSpace(app.scanner.Text())
}

func (app *TodoApp) addTask(){
	fmt.Print("Enter a task: ")
		task := app.readInput()

	if task == "" {
		fmt.Println("Task cannot be empty.")
		return
	}

	app.tasks = append(app.tasks, task)
	fmt.Println("Task added!")
}

func (app *TodoApp) viewTasks(){
	if len(app.tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	fmt.Println("Your tasks:")
	for idx, task := range app.tasks {
		fmt.Printf("%d. %s\n", idx+1, task)
	}
}

func (app *TodoApp) deleteTask() {
	if len(app.tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	fmt.Println("Select task number to delete:")
	app.viewTasks()
	fmt.Print("Enter number: ")
	input := app.readInput()

	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index < 1 || index > len(app.tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	app.tasks = append(app.tasks[:index-1], app.tasks[index:]...)
	fmt.Println("Task deleted!")
}
