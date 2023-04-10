package main

import "fmt"

func main() {
	fmt.Print("Ingrese un día de semana (1-7):")
	var diaSemana int
	fmt.Scanln(&diaSemana)
	switch diaSemana {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	}
}
