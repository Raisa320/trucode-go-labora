package main

import "fmt"

func main() {
	var rest1Nombre string = "Norte"
	rest1Calificacion := 4.5

	rest2Nombre := "Sur"
	var rest2Calificacion float32 = 3.2

	rest3Nombre := "Oeste"
	var rest3Calificacion float32 = 2.1

	fmt.Printf("El restaurante %v tiene una calificación de %v.\n", rest1Nombre, rest1Calificacion)
	fmt.Printf("El restaurante %v tiene una calificación de %v.\n", rest2Nombre, rest2Calificacion)
	fmt.Printf("El restaurante %v tiene una calificación de %v.\n", rest3Nombre, rest3Calificacion)
}
