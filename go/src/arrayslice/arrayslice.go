package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	var arr [5]string
	arr[0] = "position 1"
	fmt.Println(arr)

	arr2 := [5]int{0, 1, 3, 4, 5}
	// arr2[5] = 2
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr3)

	slice := []int{10, 2032, 23123, 9493, 30032, 03012, 0300, 332342, 2321321}

	fmt.Println(slice)
	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(reflect.TypeOf(arr3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := arr2[1:5]
	fmt.Println(slice2)
	arr2[1] = 34
	fmt.Println(slice2)

	// internal arrays
	fmt.Println("------")
	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

}
