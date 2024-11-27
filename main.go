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
		var item string
		fmt.Print("Enter an item to add to the shopping list (or type 'quit' to exit): ")
		fmt.Scanln(&item) // Get user input

		// Trim spaces and convert input to lowercase
		item = strings.ToLower(strings.TrimSpace(item))

		// Check if input is in exit_commands
		if contains(exit_commands, item) {
			fmt.Println("Exiting... Here's your shopping list:")
			displayShoppingList(shoppingList) // Display the shopping list beautifully
			break
		}

		if item == "show" {
			fmt.Println("Here's your shopping list so far:")
			displayShoppingList(shoppingList)
			continue
		}

		// Append the input to the shopping list
		shoppingList = append(shoppingList, item)
		fmt.Printf("Added '%s' to the shopping list.\n", item)
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
