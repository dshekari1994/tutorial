package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name    string
	Friends []string
}

func greet(name string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Hello, %s!", name) //send data to the channel
}

func main() {
	ch := make(chan string)

	go greet("Alice", ch)
	go greet("Bob", ch)

	// Receive from channel:
	msg1 := <-ch
	msg2 := <-ch
	fmt.Println(msg1)
	fmt.Println(msg2)

	//Create a slice of people
	people := []Person{
		{Name: "Alice", Friends: []string{"Bob", "Charlie"}},
		{Name: "David", Friends: []string{"Eve", "Frank"}},
	}

	//Print each person's friends
	for _, person := range people {
		fmt.Printf("%d %s's friends:\n", person.Name)
		for _, friend := range person.Friends {
			fmt.Printf("- %s\n", friend)
		}
	}

}
