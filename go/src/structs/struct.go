package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("file structs")
	var u user
	u.name = "Nico"
	u.age = 6
	fmt.Println(u)

	addressExemple := address{"street cool", 33}

	user2 := user{"Nico", 7, addressExemple}
	fmt.Println(user2)

	user3 := user{name: "Nico", age: 2}
	fmt.Println(user3)
}
