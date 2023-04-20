package main

import "fmt"

func multiplyMatrices(a, b [][]int, result chan<- [][]int) {
	numRowsA, numColsA := len(a), len(a[0])
	numRowsB, numColsB := len(b), len(b[0])

	if numColsA != numRowsB {
		panic("No se pueden multiplicar las matrices")
	}

	c := make([][]int, numRowsA)
	for i := range c {
		c[i] = make([]int, numColsB)
	}

	for i := 0; i < numRowsA; i++ {
		for j := 0; j < numColsB; j++ {
			for k := 0; k < numColsA; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	result <- c
}

func main() {
	a := [][]int{
		{2, 2, 5},
		{4, 6, 3},
	}

	b := [][]int{
		{1, 21},
		{2, 10},
		{11, 6},
	}

	result := make(chan [][]int)

	go multiplyMatrices(a, b, result)

	c := <-result

	fmt.Println("Resultado de la multiplicación de matrices:")
	for _, row := range c {
		fmt.Println(row)
	}
}
