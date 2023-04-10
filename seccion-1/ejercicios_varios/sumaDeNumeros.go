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
	limites := []int{50, 50, 600, 800, 0}
	descomponer := []int{0, 0, 0, 0, 0}
	var aux int = x
	for i := 0; i < len(limites); i++ {
		if aux-limites[i] > 0 && limites[i] != 0 {
			aux -= limites[i]
			descomponer[i] = limites[i]
		} else {
			descomponer[i] = aux
			break
		}
	}
	fmt.Println(x, " = ", descomponer[0], " + ", descomponer[1], " + ", descomponer[2], " + ", descomponer[3], " + ", descomponer[4])
}
