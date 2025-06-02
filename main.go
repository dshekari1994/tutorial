package main

// File: main.go
// This is a simple Go program that prints "Hello, World!" to the console.

// This program serves as a basic example of a Go application structure.
// It includes the main package, imports the fmt package for formatted I/O, and defines a main function.

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func (p Person) Greet2() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

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

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//while-like loop
	j := 0
	for j < 3 {
		fmt.Println("j is", j)
		j++
	}

	message := greet("Alice")
	fmt.Println(message)

	//fixed size array
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)

	//slices (flexible, dynamic arrays)
	numbers2 := []int{10, 20, 30, 40}
	numbers2 = append(numbers2, 50)
	fmt.Println(numbers2)

	//created a map variable and iterating over it
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	capitals["Germany"] = "Berlin"
	capitals["Russia"] = "Moscow"

	fmt.Println(capitals)

	delete(capitals, "France")

	for country, city := range capitals {
		fmt.Printf("%s: %s\n", country, city)
	}

	var numbers3 [5]int // an array of 5 integers

	//initialize
	numbers3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3[2])

	//declare and initialize a slice:
	numbers4 := []int{1, 2, 3}

	//append to a slice:
	numbers4 = append(numbers4, 4, 5)

	//iterate over a slice:
	for i, val := range numbers4 {
		fmt.Printf("Index %d: %d\n", i, val)
	}

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, "is", p.Age, "years old.")

	p.Greet2()

}
