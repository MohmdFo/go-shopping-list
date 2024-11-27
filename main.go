package main

import (
	"fmt"
	"strings"
)

// Defining a constant-like list
var exit_commands = []string{"quit", "q", "ex", "exit"}

var shoppingList []string

func main() {
	for {
		var input string
		fmt.Print("Enter an item to add to the shopping list (or type 'quit' to exit): ")
		fmt.Scanln(&input) // Get user input

		// Trim spaces and convert input to lowercase
		input = strings.ToLower(strings.TrimSpace(input))

		// Check if input is in exit_commands
		if contains(exit_commands, input) {
			fmt.Println("Exiting... Here's your shopping list:")
			displayShoppingList(shoppingList) // Display the shopping list beautifully
			break
		}

		// Append the input to the shopping list
		shoppingList = append(shoppingList, input)
		fmt.Printf("Added '%s' to the shopping list.\n", input)
	}
}

// Helper function to check if a string is in a slice
func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// Function to display the shopping list in a beautiful format
func displayShoppingList(list []string) {
	for _, item := range list {
		fmt.Printf("--> %s\n", item)
	}
}
