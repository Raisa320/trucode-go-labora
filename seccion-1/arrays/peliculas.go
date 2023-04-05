package main

import "fmt"

func main() {
	peliculas := [10]string{
		"Batman: Mask of the Phantasm",
		"Batman vs. Dracula",
		"Teen Titans: Trouble in Tokyo",
		"Superman: Doomsday",
		"Justice League: The New Frontier",
		"Wonder Woman",
		"Green Lantern",
		"Under the Red Hood",
		"The Flashpoint Paradox",
		"Son of Batman",
	}
	fmt.Println("Peliculas:", peliculas)
	fmt.Println("Segundo elemento:", peliculas[1])
	fmt.Println("Longitud:", len(peliculas))
}
