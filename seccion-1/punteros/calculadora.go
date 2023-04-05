package main

import "fmt"

func main() {
	number1 := readNumer("primer")
	number2 := readNumer("segundo")
	suma, resta, multiplicacion, division := calcular(number1, number2)
	fmt.Println("Suma: ", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicación: ", multiplicacion)
	fmt.Printf("División: %.2f\n ", division)

}

func readNumer(text string) int {
	fmt.Print("Ingrese " + text + " número: ")
	var number int
	fmt.Scanln(&number)
	return number
}

func calcular(num1, num2 int) (int, int, int, float32) {
	return num1 + num2, num1 - num2, num1 * num2, float32(num1) / float32(num2)
}
