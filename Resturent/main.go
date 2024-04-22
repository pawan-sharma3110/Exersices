package main

import (
	"fmt"
)

type menuItem struct {
	name  string
	price float64
}

func addItem(menu map[string]menuItem, name string, price float64) {
	menu[name] = menuItem{name: name, price: price}
	fmt.Println("Item added to the menu:", name)
}

func removeItem(menu map[string]menuItem, name string) {
	delete(menu, name)
	fmt.Println("Item removed from the menu:", name)
}

func printMenu(menu map[string]menuItem) {
	fmt.Println("Restaurant Menu:")
	for name, item := range menu {
		fmt.Printf("%s: $%.2f\n", name, item.price)
	}
}

func findItem(menu map[string]menuItem, name string) {
	item, exists := menu[name]
	if exists {
		fmt.Printf("Item found: %s, Price: $%.2f\n", item.name, item.price)
	} else {
		fmt.Println("Item not found in the menu.")
	}
}

func main() {
	menu := make(map[string]menuItem)

	var choice int
	for {
		fmt.Println("\nMenu Actions:")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Print Menu")
		fmt.Println("4. Find Item")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			var price float64
			fmt.Print("Enter item name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter item price: ")
			fmt.Scanln(&price)
			addItem(menu, name, price)
		case 2:
			var name string
			fmt.Print("Enter item name to remove: ")
			fmt.Scanln(&name)
			removeItem(menu, name)
		case 3:
			printMenu(menu)
		case 4:
			var name string
			fmt.Print("Enter item name to search: ")
			fmt.Scanln(&name)
			findItem(menu, name)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
