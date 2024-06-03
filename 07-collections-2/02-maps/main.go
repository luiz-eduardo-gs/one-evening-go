package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	fmt.Println("Creating user", user)
	Stats["create"]++
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)
	Stats["update"]++
}

func PurgeStats() {
	delete(Stats, "create")
	delete(Stats, "update")
}
