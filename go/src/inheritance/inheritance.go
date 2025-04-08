package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"jhon", "pedro", 21, 178}
	e1 := student{p1, "engineer", "USP"}

	fmt.Println(e1)
}
