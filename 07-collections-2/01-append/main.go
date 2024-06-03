package main

var users []string

func AddUser(u string) {
	users = append(users, u)
}

func main() {
	AddUser("Alice")
	AddUser("Bob")
}
