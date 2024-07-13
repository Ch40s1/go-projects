package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialize the status variable
	status := true
	taskList := []string{"mow lawn"}
	// Use a for loop to simulate a while loop
	for status {
		fmt.Println("1. View Tasks")
		fmt.Println("2. Create Task")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		var userAnswer int
		fmt.Scanln(&userAnswer)

		switch userAnswer {
		case 1:
			viewTasks(taskList)
		case 2:
			taskList = createTask(taskList)
		case 3:
			updateTask(taskList)
		// case 4:
		// 	deleteTask(taskList)
		case 5:
			status = false
		default:
			fmt.Println("Invalid input. Choose number between 1 and 5")
		}
	}

	fmt.Println("Exiting the program. Goodbye!")
}

func viewTasks(list []string) {
	if len(list) == 0 {
		fmt.Println("No items in list")
	} else {
		for i, item := range list {
			fmt.Printf("%d: %s\n", i+1, item)
		}
	}
}

func createTask(list []string) []string {
	fmt.Println("Enter the new task:")
	userData, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	results := strings.TrimSuffix(userData, "\n")
	return append(list, results)
}
func updateTask(list []string) {
	viewTasks(list)
	fmt.Println("Which task would you like to modify?enter number of task.")
	var userData int
	fmt.Scanln(&userData)
	if userData < 0 || userData < len(list) {
		fmt.Println("please enter number of a task.")
	} else {
		fmt.Println("Enter update:")
		modify, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		results := strings.TrimSuffix(modify, "\n")
		list[userData-1] = results
	}
}

// func deleteTask(list []string) {
// 	//pass
// }
