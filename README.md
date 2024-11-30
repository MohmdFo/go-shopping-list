# Shopping List CLI Application

This is a simple terminal-based **shopping list application** written in Go. It allows users to manage a shopping list interactively, providing commands to add, remove, search, display items, view the total number of items, and exit the program. The application includes clear terminal functionality for a more user-friendly experience.

---

## Features

- **Add Items:** Add new items to your shopping list.
- **Display List:** View all items in your shopping list.
- **Remove Items:** Remove specific items from your list.
- **Search Items:** Search for specific items to check if they exist in the list.
- **View Total:** Display the total number of items in your shopping list.
- **Help:** View a list of available commands and their descriptions.
- **Exit:** Quit the application and display the final shopping list.
- **Clear Screen:** Automatically clears the terminal for a clean user interface.

---

## Commands

- **`quit`** or **`q`**: Exit the program and display the final shopping list.
- **`help`**: Show a list of available commands and their descriptions.
- **`show`**: Display the current shopping list.
- **`remove`**: Remove a specified item from the shopping list.
- **`search`**: Check if a specific item exists in the shopping list.
- **`total`**: Display the total number of items in the shopping list.
- **`priority`**: Change the priority of an item by moving it to a different position in the list.
- **Adding Items**: Simply type the name of an item to add it to your list.

---

## How to Run

1. **Install Go:** Ensure that [Go](https://golang.org/dl/) is installed on your system.
2. **Download the Code:**
   ```bash
   git clone https://github.com/your-username/shopping-list-cli.git
   cd shopping-list-cli
   ```
3. **Run the Application:**
   ```bash
   go run main.go
   ```

---

## How to Use

### 1. Adding Items
Simply type the name of an item to add it to the list:
```text
Enter an item (type 'quit' to exit, 'help' for commands): apple
Added 'apple' to the shopping list.
```

### 2. Viewing the Shopping List
Type `show` to view all items in your list:
```text
Enter an item (type 'quit' to exit, 'help' for commands): show
Here's your shopping list so far:
--> apple
--> banana
```

### 3. Viewing the Total Number of Items
Type `total` to display the total number of items in the shopping list:
```text
Enter an item (type 'quit' to exit, 'help' for commands): total
There are a total of 2 items in your shopping list.
```

### 4. Removing Items
Type `remove`, then specify the item to remove:
```text
Enter an item (type 'quit' to exit, 'help' for commands): remove
Enter the item you want to remove: banana
Item 'banana' has been removed from the shopping list.
```

### 5. Searching for Items
Type `search`, then enter the item to search for:
```text
Enter an item (type 'quit' to exit, 'help' for commands): search
Enter the item you want to search: apple
Item 'apple' exists in the list.
```

### 6. Getting Help
Type `help` to view the list of available commands:
```text
Enter an item (type 'quit' to exit, 'help' for commands): help

Available commands:
  quit    - Exit the program
  help    - Show this help message
  show    - Display the shopping list
  total   - Display the total number of items in the shopping list
  remove  - Remove an item from the shopping list
  search  - Search for an item in the shopping list

Simply enter an item to add it to the shopping list.
```

### 7. Exiting the Program
Type `quit` or `q` to exit the application and display the final shopping list:
```text
Enter an item (type 'quit' to exit, 'help' for commands): quit
Exiting... Here's your shopping list:
--> apple
--> orange
```

---

## Code Structure

1. **`main` Function:**
   - The main loop for handling user input and commands.
   - Handles adding items, displaying help, showing the list, and other operations.

2. **`clearScreen` Function:**
   - Clears the terminal for a clean user interface.
   - Supports both Linux/macOS (`clear`) and Windows (`cls`).

3. **Command-Specific Functions:**
   - `displayHelp`: Displays a list of available commands.
   - `displayShoppingList`: Prints the current shopping list in a formatted style.
   - `removeItemFromList`: Removes an item from the list.
   - `search`: Checks if an item exists in the shopping list.
   - `displayTotal`: Displays the total number of items in the shopping list.

4. **Utility Functions:**
   - `contains`: Checks if a slice contains a specific item.

---

## Example Run

```text
Enter an item (type 'quit' to exit, 'help' for commands): apple
Added 'apple' to the shopping list.

Enter an item (type 'quit' to exit, 'help' for commands): banana
Added 'banana' to the shopping list.

Enter an item (type 'quit' to exit, 'help' for commands): show
Here's your shopping list so far:
--> apple
--> banana

Enter an item (type 'quit' to exit, 'help' for commands): total
There are a total of 2 items in your shopping list.

Enter an item (type 'quit' to exit, 'help' for commands): search
Enter the item you want to search: apple
Item 'apple' exists in the list.

Enter an item (type 'quit' to exit, 'help' for commands): remove
Enter the item you want to remove: banana
Item 'banana' has been removed from the shopping list.

Enter an item (type 'quit' to exit, 'help' for commands): quit
Exiting... Here's your shopping list:
--> apple
```

---

## Contribution

Feel free to open issues or submit pull requests to improve this application. Contributions are welcome!

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
