package main

import "fmt"

var count = 0

func AllocateBuffer() *string {
	if count++; count > 3 {
		return nil
	}

	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
