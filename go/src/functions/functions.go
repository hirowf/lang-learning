package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}
func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum := sum(10, 20)
	fmt.Println(sum)

	var f = func(txt string) string {
		//fmt.Println(txt)
		return txt
	}

	res := f("function text")
	fmt.Println(res)

	resMathCalsSum, _ := mathCalcs(23, 20)
	fmt.Println(resMathCalsSum)
}
