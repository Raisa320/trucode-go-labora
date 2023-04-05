package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono string
}

func cambiarCiudad(persona *Persona, ciudad string) bool {
	if ciudad != persona.ciudad {
		persona.ciudad = ciudad
		return true
	}
	return false
}

func main() {
	persona1 := Persona{"Juan", 30, "Madrid", "555-1234"}
	persona2 := Persona{"Ana", 25, "Barcelona", "555-5678"}
	fmt.Println("Persona 1", persona1)
	fmt.Println("Persona 2", persona2)

	condicion := cambiarCiudad(&persona1, "Valencia")
	printPersonalizado(persona1, condicion)
	condicion = cambiarCiudad(&persona2, "Barcelona")
	printPersonalizado(persona1, condicion)

}

func printPersonalizado(p Persona, condicion bool) {
	if condicion {
		fmt.Println("Persona con ciudad actualizada: ", p)
	} else {
		fmt.Println("Persona: ", p)
	}
}
