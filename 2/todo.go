package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){


	tasks := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo App ---")
		fmt.Println("1. Add a task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")
		fmt.Println("Choose an option:")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Enter task: ")
			scanner.Scan();
			task := scanner.Text();
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
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid Option, choose again.")
		}
	}
}