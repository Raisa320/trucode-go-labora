package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/raisa320/trucode-go-labora/seccion-4/interfaz-clase/figthers"
)

func main() {
	var police figthers.Police = figthers.Police{BaseFigther: figthers.BaseFigther{Life: 20}, Armour: 5}
	var criminal figthers.Criminal = figthers.Criminal{BaseFigther: figthers.BaseFigther{Life: 22}}
	var contenders []figthers.Contender = make([]figthers.Contender, 2)

	randomValueBetweenOneAndZero := rand.Intn(2) //0-1
	contenders[randomValueBetweenOneAndZero] = &police
	contenders[(randomValueBetweenOneAndZero+1)%2] = &criminal

	var areBothAlive = police.IsAlive() && criminal.IsAlive()

	for areBothAlive {
		intensity := contenders[0].ThrowAttack()
		fmt.Println(contenders[0].GetName(), "tira golpe con intensidad  =", intensity)
		contenders[1].RecieveAttack(intensity)
		if contenders[1].IsAlive() {
			instensityy := contenders[1].ThrowAttack()
			fmt.Println(contenders[1].GetName(), "tira golpe con intensidad  =", instensityy)
			contenders[0].RecieveAttack(instensityy)
		}
		fmt.Printf("PoliceLife=%d, CriminalLife=%d\n", police.Life, criminal.Life)
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(time.Second * 3)
	}
}

func main_legacy() {
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
