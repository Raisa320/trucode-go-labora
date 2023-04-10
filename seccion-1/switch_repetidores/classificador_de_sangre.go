package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Ingrese tipo de sangre:")
	var tipoSangre string
	fmt.Scanln(&tipoSangre)
	fmt.Println(clasificarSangre(tipoSangre))
}

func clasificarSangre(tipoSangre string) string {
	cadenaTipoSangre := "Grupo sanguíneo "
	cadenaTipoSangre += strings.ToUpper(tipoSangre[:len(tipoSangre)-1])
	switch strings.ToUpper(tipoSangre) {
	case "A+", "B+", "AB+", "O+":
		cadenaTipoSangre += ", factor Rh positivo"
	case "A-", "B-", "AB-", "O-":
		cadenaTipoSangre += ", factor Rh negativo"
	}
	return cadenaTipoSangre
}
