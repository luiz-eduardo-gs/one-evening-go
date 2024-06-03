package main

import "fmt"

func Greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	Greet("Alice")
	Greet("Bob")
}
