package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptrA *int = &a
	fmt.Println("Valores iniciales: a=", a, ", b = ", b)
	*ptrA, b = b, *ptrA
	fmt.Println("Valores intercambiados: a=", a, ", b = ", b)
}
