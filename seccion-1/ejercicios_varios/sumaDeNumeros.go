/*
Enuciado: Expresar a un número cualquiera como la suma de una serie de números siguiendo
las restricciones impuestas a continuación
Sea x el número.
x = s1 + s2 + s3 + s4 + s5

Donde
0 < s1 < 50
0 < s2 < 50
0 < s3 < 600
0 < s4 < 800
0 < s5 < (Infinito)
*/
package main

import "fmt"

func main() {
	fmt.Print("Ingrese un número: ")
	var x int
	fmt.Scanln(&x)
	descomponer := []int{0, 0, 0, 0, 0}
	segmentarValorPorRangos(x, descomponer)

	fmt.Println(x, " = ", descomponer[0], " + ", descomponer[1], " + ", descomponer[2], " + ", descomponer[3], " + ", descomponer[4])
}

func segmentarValorPorRangos(x int, descomponer []int) {
	limites := []int{50, 50, 600, 800, 0}
	for i := 0; i < len(limites); i++ {
		if x-limites[i] > 0 && limites[i] != 0 {
			x -= limites[i]
			descomponer[i] = limites[i]
		} else {
			descomponer[i] = x
			return
		}
	}
}
