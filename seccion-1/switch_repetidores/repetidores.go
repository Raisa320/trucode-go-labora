package main

import "fmt"

type Planeta struct {
	nombre    string
	cntdLunas int
}

func main() {
	planetas := []Planeta{
		{"Mercurio", 0},
		{"Venus", 0},
		{"Tierra", 1},
		{"Marte", 2},
		{"Júpiter", 79},
		{"Saturno", 62},
		{"Urano", 27},
		{"Neptuno", 14},
		{"Plutón", 6}}

	for _, planeta := range planetas {
		fmt.Print(planeta.nombre)
		if planeta.cntdLunas == 0 {
			fmt.Println(" no tiene lunas.")
		} else if planeta.cntdLunas == 1 {
			fmt.Println(" tiene", planeta.cntdLunas, "luna.")
		} else {
			fmt.Println(" tiene", planeta.cntdLunas, "lunas.")
		}

	}
}
