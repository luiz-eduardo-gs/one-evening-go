package main

import (
	"fmt"
	"os"
)

func CheckFile(name string) bool {
	_, err := os.ReadFile(name)

	if err != nil {
		return false
	}

	return true
}

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}
