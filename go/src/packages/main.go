package main

import (
	"fmt"
	"pkgmod/auxfolder"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("writing from main file")
	auxfolder.Write()

	error := checkmail.ValidateFormat("")

	fmt.Println(error)
}
