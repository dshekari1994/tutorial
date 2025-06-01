package main

// File: main.go
// This is a simple Go program that prints "Hello, World!" to the console.

// This program serves as a basic example of a Go application structure.
// It includes the main package, imports the fmt package for formatted I/O, and defines a main function.

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// The main function is the entry point of the program.
	// It uses the fmt package to print a message to the console.
	// This code also demonstrates variable declaration and initialization in Go.

	var name string = "Alice"
	var age int = 25
	height := 5.5 //shorthand declaration

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	if age >= 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is not an adult.")
	}
}
