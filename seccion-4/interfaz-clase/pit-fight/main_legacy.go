package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/raisa320/trucode-go-labora/seccion-4/interfaz-clase/figthers"
)

func Main_legacy() {
	var police figthers.Police = figthers.Police{BaseFigther: figthers.BaseFigther{Life: 20}, Armour: 5}
	var criminal figthers.Criminal = figthers.Criminal{BaseFigther: figthers.BaseFigther{Life: 20}}

	randomValueBetweenOneAndZero := rand.Intn(2)
	policeHitFirst := randomValueBetweenOneAndZero == 1

	var areBothAlive = police.IsAlive() && criminal.IsAlive()

	for areBothAlive {
		if policeHitFirst {
			intensity := police.ThrowAttack()
			fmt.Println("Policia tira golpe con intensidad=", intensity)
			criminal.RecieveAttack(intensity)
			if criminal.IsAlive() {
				intensity := criminal.ThrowAttack()
				fmt.Println("Criminal tira golpe con intensidad=", intensity)
				police.RecieveAttack(intensity)
			}
		} else {
			intensity := criminal.ThrowAttack()
			fmt.Println("Criminal tira golpe con intensidad=", intensity)
			police.RecieveAttack(intensity)
			if police.IsAlive() {
				intensity := police.ThrowAttack()
				fmt.Println("Policia tira golpe con intensidad=", intensity)
				criminal.RecieveAttack(intensity)
			}
		}
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(time.Second * 3)
		fmt.Printf("PoliceLife=%d, CriminalLife=%d\n", police.Life, criminal.Life)
	}

}
