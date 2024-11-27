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
		fmt.Print("Enter an item (type 'quit' to exit, 'help' for commands): ")
		fmt.Scanln(&item) // Get user input

		// Trim spaces and convert input to lowercase
		item = strings.ToLower(strings.TrimSpace(item))

		// Check if input is in exit_commands
		if contains(exit_commands, item) {
			fmt.Println("Exiting... Here's your shopping list:")
			displayShoppingList(shoppingList) // Display the shopping list beautifully
			break
		}

		// Show help if the user types 'help'
		if item == "help" {
			displayHelp()
			continue
		}

		// Check if the user wants to show the shopping list
		if item == "show" {
			fmt.Println("Here's your shopping list so far:")
			displayShoppingList(shoppingList)
			continue
		}

		if item == "search" {
			search()
			continue
		}

		// Check if the user wants to remove an item
		if item == "remove" {
			removeItemFromList() // Call function to handle item removal
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

// Function to display help information
func displayHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  quit    - Exit the program")
	fmt.Println("  help    - Show this help message")
	fmt.Println("  show    - Display the shopping list")
	fmt.Println("  remove  - Remove an item from the shopping list")
	fmt.Println("  search  - Search for an item in the shopping list")
	fmt.Println("\nSimply enter an item to add it to the shopping list.")
}

// Function to search for an item in the shopping list
func search() {
	var item_to_search string
	fmt.Print("Enter the item you want to search: ")
	fmt.Scanln(&item_to_search)                                         // Get the item to search
	item_to_search = strings.ToLower(strings.TrimSpace(item_to_search)) // Normalize the input

	found := false // Flag to track if the item is found

	for _, value := range shoppingList {
		if value == item_to_search {
			fmt.Printf("Item '%s' exists in the list.\n", item_to_search)
			found = true // Set the flag to true if the item is found
			break
		}
	}

	if !found { // Check the flag after the loop
		fmt.Printf("Item '%s' does not exist in the list.\n", item_to_search)
	}
}

// Function to display the shopping list in a beautiful format
func displayShoppingList(list []string) {
	if len(list) == 0 {
		fmt.Println("The shopping list is empty.")
		return
	}
	for _, item := range list {
		fmt.Printf("--> %s\n", item)
	}
}

// Function to remove an item from the shopping list
func removeItemFromList() {
	var itemToRemove string
	fmt.Print("Enter the item you want to remove: ")
	fmt.Scanln(&itemToRemove)                                       // Get the item to remove
	itemToRemove = strings.ToLower(strings.TrimSpace(itemToRemove)) // Normalize the input

	index := -1
	for i, value := range shoppingList { // Find the index of the item
		if value == itemToRemove {
			index = i
			break
		}
	}

	if index == -1 { // Item not found
		fmt.Printf("Item '%s' does not exist in the shopping list.\n", itemToRemove)
	} else {
		// Remove the item by slicing
		shoppingList = append(shoppingList[:index], shoppingList[index+1:]...)
		fmt.Printf("Item '%s' has been removed from the shopping list.\n", itemToRemove)
	}
}
