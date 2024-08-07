package main

import "fmt"

func main() {
	list := []rune{'s', 'a', 'l', 'm', 'a', 'n'}
	slice := list[1:]
	fmt.Println(string(slice))
}
