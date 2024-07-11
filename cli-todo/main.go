package main

import "fmt"

func main() {
	// Initialize the status variable
	status := true

	// Use a for loop to simulate a while loop
	for status {
		fmt.Println("press 1 to continue")
		fmt.Println("press 2 to exit")
		var userAnswer int
		fmt.Scanln(&userAnswer)

		if userAnswer == 2 {
			status = false
		} else if userAnswer == 1 {
			fmt.Println("hello!")
		} else {
			fmt.Println("Invalid input. Please press 1 or 2.")
		}
	}

	fmt.Println("Exiting the program. Goodbye!")
}
