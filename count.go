package main

import "fmt"

func main() {
	fmt.Print("Enter text: ")
	var mytext string
	var size int
	fmt.Scanln(&mytext)
	size = 0
	for position, character := range mytext {
		fmt.Println(position)
		fmt.Println(character)
		size++
	}
	fmt.Print("Size of text is: ")
	fmt.Println(size)
}
