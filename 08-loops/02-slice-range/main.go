package main

func Sum(x ...int) int {
	var sum int
	for _, n := range x {
		sum += n
	}

	return sum
}

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}
