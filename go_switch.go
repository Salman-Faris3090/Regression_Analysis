package main

import "fmt"

func main() {
	var input interface{} = 12.5

	switch input.(type) {
	case int:
		fmt.Println("The Entered input is a Number")
	case string:
		fmt.Println("The Entered input is a String")
	case float64:
		fmt.Println("The Entered input is a Floating Number")
	default:
		fmt.Println("Unrecognized input")
	}
}
