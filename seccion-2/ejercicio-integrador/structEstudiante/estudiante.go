package est

import (
	"fmt"
	"sort"
)

type Estudiante struct {
	Nombre string
	Nota   float64
	Codigo string
}

func ImprimirEstudiantes(estudiantes []Estudiante) {
	for i, estudiante := range estudiantes {
		fmt.Printf("===== Estudiante %d ======\n", i+1)
		fmt.Println("Nombre:", estudiante.Nombre)
		fmt.Println("Nota:", estudiante.Nota)
		fmt.Println("Código:", estudiante.Codigo)
	}
}

func BuscarPorCodigo(estudiantes []Estudiante) {
	var cod string
	fmt.Print("Ingresa el código del estudiante: ")
	fmt.Scan(&cod)
	for _, estudiante := range estudiantes {
		if estudiante.Codigo == cod {
			fmt.Printf("\nCódigo: %s - Nombre: %s, Nota: %.2f\n", estudiante.Codigo, estudiante.Nombre, estudiante.Nota)
			break
		}
	}
}

func CrearEstudiante(estudiantes *[]Estudiante) {
	fmt.Println("======= Crear Estudiante =========")
	var estudiante Estudiante
	fmt.Println("Nombre: ")
	fmt.Scan(&estudiante.Nombre)

	fmt.Println("Nota: ")
	fmt.Scan(&estudiante.Nota)

	fmt.Println("Código: ")
	fmt.Scan(&estudiante.Codigo)

	*estudiantes = append(*estudiantes, estudiante)
}

func OrdenarEstudiantesA(criterio string, estudiantes []Estudiante) {
	switch criterio {
	case "1":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Nombre < estudiantes[j].Nombre
		})
	case "2":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Nota < estudiantes[j].Nota
		})
	case "3":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Codigo < estudiantes[j].Codigo
		})

	default:
		fmt.Print("no se cumplio ninguno")
	}

}

func OrdenarEstudiantesD(criterio string, estudiantes []Estudiante) {
	switch criterio {
	case "1":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Nombre > estudiantes[j].Nombre
		})
	case "2":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Nota > estudiantes[j].Nota
		})
	case "3":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Codigo > estudiantes[j].Codigo
		})
	default:
		fmt.Print("no se cumplio ninguno")
	}

}
