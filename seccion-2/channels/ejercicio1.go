// Ejercicio 1: Suma de números en un rango utilizando gorutinas y canales
// Escribe un programa en Go que sume todos los números en un rango dado (por ejemplo, de 1 a 100)
// utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas.

package main

import "fmt"

func main() {
	init := 1
	end := 100
	mitad := end / 2

	resultCh := make(chan int)
	go suma(init, mitad, resultCh)
	go suma(mitad+1, end, resultCh)

	// Espera a que ambas gorutinas terminen y obtiene los resultados
	sum1 := <-resultCh
	sum2 := <-resultCh

	// Imprime la suma total
	fmt.Println("Suma total:", sum1+sum2)
}

func suma(init, end int, resultCh chan int) {
	suma := 0
	for i := init; i <= end; i++ {
		suma += i
	}
	resultCh <- suma
}
