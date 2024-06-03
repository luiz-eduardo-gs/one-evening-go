package main

import (
	"fmt"
)

func WordGenerator(words []string) func() string {
	i := 0

	return func() string {
		if i > len(words)-1 {
			i = 0
		}

		word := words[i]
		i++
		return word
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
