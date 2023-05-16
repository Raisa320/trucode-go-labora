package main

import (
	"fmt"

	"github.com/raisa320/trucode-go-labora/seccion-3/TESTING/cadenasRotation/slices"
)

func main() {
	cadena := "ABCDF"
	cadena2 := "ABCDE"
	for i := 0; i < 3; i++ {
		cadena = slices.RotateRightMyVersion(cadena)
	}
	fmt.Println(slices.RotateRightMyVersion(cadena2))
	fmt.Println(slices.RotateRightVituVersion(cadena2))
	fmt.Println(cadena)
}
