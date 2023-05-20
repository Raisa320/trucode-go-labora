package figthers

import "math/rand"

type Police struct {
	BaseFigther //promocion = promueve  los miembros  aqui? implicitamente
	Armour      int
}

func (p Police) ThrowAttack() int {
	return rand.Intn(10)
}

func (p *Police) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := p.Armour - intensity
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= intensity
			intensity = 0
		} else {
			p.Armour = 0
			intensity = -diff
		}
	}
	p.Life -= intensity
}

func (p Police) GetName() string {
	return "Policia"
}
