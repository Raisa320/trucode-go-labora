package main

import "fmt"

func main() {
	var equipo1, equipo2, equipo3, equipo4, equipo5 string = "San Juan", "San Pedro", "Ten", "Eleven", "Triple"

	var equipo1Titulos, equipo2Titulos, equipo3Titulos, equipo4Titulos, equipo5Titulos int = 12, 21, 11, 10, 15

	fmt.Printf("El equipo %v ha ganado %v títulos.\n", equipo1, equipo1Titulos)
	fmt.Printf("El equipo %v ha ganado %v títulos.\n", equipo2, equipo2Titulos)
	fmt.Printf("El equipo %v ha ganado %v títulos.\n", equipo3, equipo3Titulos)
	fmt.Printf("El equipo %v ha ganado %v títulos.\n", equipo4, equipo4Titulos)
	fmt.Printf("El equipo %v ha ganado %v títulos.\n", equipo5, equipo5Titulos)
}
