package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){


	tasks := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo App ---")
		fmt.Println("1. Add a task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Tasks")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option:")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if choice == "" {
			fmt.Println("Please enter a valid option.")
			continue
		}

		switch choice {
		case "1":
			fmt.Println("Enter task: ")
			scanner.Scan();
			task := strings.TrimSpace(scanner.Text());

			if task == "" {
				fmt.Println("Task cannot be empty.")
				continue
			}

			tasks = append(tasks, task)
			fmt.Println("Task added successfully.")

		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			}else{
				fmt.Println("Your tasks:")
				for idx,task := range tasks {
					fmt.Printf("%d. %s\n", idx+1,task)
				}
			}

		case "3":
			if len(tasks) == 0 {
				fmt.Println("No tasks available to delete.")
				continue
			}

			fmt.Println("Select a task to delete: ")
			for idx,task := range tasks {
					fmt.Printf("%d. %s\n", idx+1,task)
				}

			fmt.Println("Enter number:")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			var index int
			_, err := fmt.Sscanf(input, "%d", &index)
			if err != nil || index < 1 || index > len(tasks){
				fmt.Println("Invalid task tumber");
				continue
			}

			tasks = append(tasks[:index-1], tasks[index:]...)
			fmt.Println("Task deleted!")

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid Option, choose again.")
		}
	}
}