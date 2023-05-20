package figthers

import "math/rand"

type Paladin struct {
	BaseFigther
}

func (p Paladin) ThrowAttack() int {
	attackIntensity := rand.Intn(15)
	if p.Life <= p.LifeInitial/2 {
		return attackIntensity / 2
	}
	return attackIntensity
}

func (p *Paladin) RecieveAttack(intensity int) {
	p.Life -= intensity
}

func (p Paladin) GetName() string {
	return "Paladín"
}
