package main

import (
	"fmt"
	est "moduloIntegrador/structEstudiante"
	"moduloIntegrador/utils"
	"sort"
)

func menuOrdenar(opcion string, estudiantes []est.Estudiante) {
	opMenu := []string{"ORDENAR", "ordenar por Nombre", "ordenar por Nota", "ordenar por Codigo"}
	switch opcion {
	case "1":
		utils.ImprimirMenu(opMenu)
		fmt.Scan(&opcion)
		est.OrdenarEstudiantesA(opcion, estudiantes)
	case "2":
		utils.ImprimirMenu(opMenu)
		fmt.Scan(&opcion)
		est.OrdenarEstudiantesD(opcion, estudiantes)
	}
}
func notaPromedio(estudiantes *[]est.Estudiante) {
	var suma float64 = 0
	var promedio float64 = 0

	for _, estudiante := range *estudiantes {
		suma += estudiante.Nota
	}

	if len(*estudiantes) != 0 {
		promedio = suma / float64(len(*estudiantes))
	}

	fmt.Printf("Nota promedio de los estudiantes: %.2f\n", promedio)
}

func buscarNotaMaxMin(estudiantes []est.Estudiante) {
	sort.Slice(estudiantes, func(i, j int) bool {
		return estudiantes[i].Nota > estudiantes[j].Nota
	})
	fmt.Printf("El estudiante %v tiene la mayor nota: %.2f\n", estudiantes[0].Nombre, estudiantes[0].Nota)
	fmt.Printf("El estudiante %v tiene la menor nota: %.2f\n", estudiantes[len(estudiantes)-1].Nombre, estudiantes[len(estudiantes)-1].Nota)
}

func main() {
	var estudiantes []est.Estudiante

	defer func() {
		fmt.Println("==== Estadisticas ====")
		notaPromedio(&estudiantes)
		buscarNotaMaxMin(estudiantes)
	}()

	var opcion string
	opcionesPrincipal := []string{"GESTION ESTUDIANTES", "Crear estudiantes", "Ordenar estudiantes", "Buscar estudiantes", "listar estudiantes", "Salir"}

	for {
		utils.ImprimirMenu(opcionesPrincipal)
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			est.CrearEstudiante(&estudiantes)
		case "2":
			opOrdenar := []string{"ORDENAR", "Ordenar en forma ascendente", "Ordenar en forma descendente"}
			utils.ImprimirMenu(opOrdenar)
			fmt.Scan(&opcion)
			menuOrdenar(opcion, estudiantes)
			fmt.Println("lista ordenada")
			est.ImprimirEstudiantes(estudiantes)
		case "3":
			est.BuscarPorCodigo(estudiantes)
		case "4":
			est.ImprimirEstudiantes(estudiantes)
		case "5":
			return
		default:
			fmt.Print("la opcion ingresada no es valida")
		}
	}

}
