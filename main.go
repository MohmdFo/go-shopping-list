package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
			clearScreen()
			fmt.Println("Exiting... Here's your shopping list:")
			displayShoppingList(shoppingList) // Display the shopping list beautifully
			break
		}

		// Show help if the user types 'help'
		if item == "help" {
			clearScreen()
			displayHelp()
			continue
		}

		// Check if the user wants to show the shopping list
		if item == "show" {
			clearScreen()
			fmt.Println("Here's your shopping list so far:")
			displayShoppingList(shoppingList)
			continue
		}

		// Check if the user wants to see the total number of items
		if item == "total" {
			clearScreen()
			displayTotal()
			continue
		}

		// Check if the user wants to search for an item
		if item == "search" {
			clearScreen()
			search()
			continue
		}

		// Check if the user wants to remove an item
		if item == "remove" {
			clearScreen()
			removeItemFromList() // Call function to handle item removal
			continue
		}

		// Check if the user wants to change the priority of an item
		if item == "priority" {
			clearScreen()
			changePriority() // Call function to change item priority
			continue
		}

		// Check for duplicate items
		if contains(shoppingList, item) {
			fmt.Printf("Item '%s' is already in the shopping list. Please add something else.\n", item)
			continue
		}

		// Append the input to the shopping list
		shoppingList = append(shoppingList, item)
		clearScreen()
		fmt.Printf("Added '%s' to the shopping list.\n", item)
	}
}

// Function to clear the terminal screen
func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	if strings.Contains(os.Getenv("OS"), "Windows") {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
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
	fmt.Println("  quit      - Exit the program")
	fmt.Println("  help      - Show this help message")
	fmt.Println("  show      - Display the shopping list")
	fmt.Println("  total     - Display the total number of items in the shopping list")
	fmt.Println("  remove    - Remove an item from the shopping list")
	fmt.Println("  search    - Search for an item in the shopping list")
	fmt.Println("  priority  - Change the priority of an item in the shopping list")
	fmt.Println("\nSimply enter an item to add it to the shopping list.")
}

// Function to display the total number of items in the shopping list
func displayTotal() {
	totalItems := len(shoppingList)
	fmt.Printf("There are a total of %d items in your shopping list.\n", totalItems)
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

// Function to change the priority of an item in the shopping list
func changePriority() {
	var item string
	fmt.Print("Enter the item whose priority you want to change: ")
	fmt.Scanln(&item)
	item = strings.ToLower(strings.TrimSpace(item))

	index := -1
	for i, value := range shoppingList { // Find the index of the item
		if value == item {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Item '%s' does not exist in the shopping list.\n", item)
		return
	}

	var newPositionStr string
	fmt.Printf("Enter the new position for '%s' (1 to %d): ", item, len(shoppingList))
	fmt.Scanln(&newPositionStr)

	newPosition, err := strconv.Atoi(newPositionStr)
	if err != nil || newPosition < 1 || newPosition > len(shoppingList) {
		fmt.Println("Invalid position. Please enter a valid number.")
		return
	}

	// Adjust for zero-based indexing
	newPosition--

	// Remove the item from the current position
	shoppingList = append(shoppingList[:index], shoppingList[index+1:]...)

	// Insert the item at the new position
	shoppingList = append(shoppingList[:newPosition], append([]string{item}, shoppingList[newPosition:]...)...)

	fmt.Printf("Item '%s' has been moved to position %d.\n", item, newPosition+1)
}
