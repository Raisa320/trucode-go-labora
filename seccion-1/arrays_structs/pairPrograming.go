package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	telefono int64
}

func (p *Persona) incrementarEdad() {
	p.edad++
}

func (p *Persona) actualizarTelefono(telefono int64) {
	p.telefono = telefono
}

func main() {
	persona1 := Persona{"Raisa", 23, 955551476}
	persona2 := Persona{"Vanessa", 27, 955551477}
	persona1.incrementarEdad()

	fmt.Println(persona1)
	fmt.Println(persona2)

	personas := [5]Persona{
		persona1,
		persona2,
		{"Carmen", 35, 955551455},
		{"Karla", 98, 955551499},
		{"Juana", 32, 955551477}}

	fmt.Println(buscarEdad(personas, 32))
	fmt.Println(buscarEdad(personas, 15))

	personaNueva := crearPersona("Carlos", 12, 955551478)
	personaNueva.actualizarTelefono(988881456)
	fmt.Println(*personaNueva)

}

func buscarEdad(personas [5]Persona, edad int) *Persona {
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad {
			return &personas[i]
		}
	}
	var puntero *Persona
	return puntero
}

func crearPersona(nombre string, edad int, telefono int64) *Persona {
	return &Persona{nombre, edad, telefono}
}
