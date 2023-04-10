package main

import "fmt"

func main() {
	//1030
	var segundosInput int
	fmt.Print("Ingrese tiempo en segundos: ")
	fmt.Scanln(&segundosInput)

	segundos, minutosTotal := descomponer(segundosInput, 60)
	minutos, horasTotal := descomponer(minutosTotal, 60)
	horas, dias := descomponer(horasTotal, 24)

	fmt.Println(segundosInput, "segundos", "son:", dias, "días,", horas, "horas,", minutos, "minutos con", segundos, "segundos.")

}

func descomponer(input int, tiempo int) (int, int) {
	if input > tiempo {
		return input % tiempo, input / tiempo
	} else {
		return input, input
	}
}
