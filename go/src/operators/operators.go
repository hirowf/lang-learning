package main

import "fmt"

func main() {
	// arithmetics
	sum := 1 + 2
	sub := 1 - 2
	mul := 1 * 2
	div := 19 / 4
	rest := 10 % 2

	fmt.Println(sum, sub, mul, div, rest)

	var num1 int16 = 10
	var num2 int16 = 25
	x := num1 + num2
	fmt.Println(x)

	// assigned

	var var1 string = "String"
	var2 := "String2"

	fmt.Println(var1, var2)

	// relational operators

	fmt.Println(1 > 2)
	fmt.Println(2 >= 3)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// logical operators

	// AND OR !
	fmt.Println("----------------")
	ver, fals := true, false
	fmt.Println(ver && fals)
	fmt.Println(ver || fals)
	fmt.Println(!ver)

	// operadores unarios

	number := 10
	number++
	number += 15 // number = number + 15
	fmt.Println(number)

	number--
	number -= 20 // number = number - 20
	number *= 20 // number = number * 3
	number %= 3

	fmt.Println(number)

	// operador ternario
	fmt.Println("-----operador ternario-----")

	// text := number > 5 ? "Maior que 5" : "menor que 5"
	var text string
	if number > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}
	fmt.Println(text)

}
