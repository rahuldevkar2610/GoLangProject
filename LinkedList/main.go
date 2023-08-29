package main

import "fmt"

func main() {
	fmt.Println("Node Assignment...!!!!")

	var head *Node = nil

	for {

		fmt.Println("\nMenu:")
		fmt.Println("1. Add Node")
		fmt.Println("2. Delete Node")
		fmt.Println("3. Traverse")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var value int
			fmt.Print("Enter value: ")
			fmt.Scanln(&value)
			head = AddNode(head, value)
		case 2:
			var value int
			fmt.Print("Enter value to delete: ")
			fmt.Scanln(&value)
			head = DeleteNode(head, value)
		case 3:
			Traverse(head)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
