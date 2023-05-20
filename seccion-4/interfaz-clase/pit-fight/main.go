package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/raisa320/trucode-go-labora/seccion-4/interfaz-clase/figthers"
)

func createContendersAndAssignPosition(contenders ...figthers.Contender) []figthers.Contender {
	sizeContenders := len(contenders)
	var contendersList []figthers.Contender = make([]figthers.Contender, sizeContenders)
	var randomPosition = rand.Intn(sizeContenders)
	contendersList[randomPosition] = contenders[0]
	for i := 1; i < sizeContenders; i++ {
		contendersList[(randomPosition+i)%sizeContenders] = contenders[i]
	}
	return contendersList
}

func battle(contenders []figthers.Contender, contendersDead int) int {
	for _, contender := range contenders {
		if contender.IsAlive() {
			intensity := contender.ThrowAttack()
			fmt.Println(contender.GetName(), "tira golpe con intensidad  =", intensity)
			for _, nextContender := range contenders {
				if contender != nextContender && nextContender.IsAlive() {
					nextContender.RecieveAttack(intensity)
					if nextContender.IsAlive() {
						fmt.Println(nextContender.GetName(), "le queda de vida  =", nextContender.ActualLife())
					} else {
						fmt.Println(nextContender.GetName(), "ha muerto")
						contendersDead++
					}
				}
			}
			fmt.Println()
		}
	}
	return contendersDead
}

func contenderWinner(contenders []figthers.Contender) {
	for _, contender := range contenders {
		if contender.IsAlive() {
			fmt.Printf("\n%v ha ganado la pelea con una vida restante de %v", contender.GetName(), contender.ActualLife())
			break
		}
	}
}

func main() {
	var police figthers.Police = figthers.Police{BaseFigther: figthers.BaseFigther{Life: 20}, Armour: 5}
	var criminal figthers.Criminal = figthers.Criminal{BaseFigther: figthers.BaseFigther{Life: 22}}
	var paladin figthers.Paladin = figthers.Paladin{BaseFigther: figthers.BaseFigther{Life: 22, LifeInitial: 22}}

	var contenders = createContendersAndAssignPosition(&police, &criminal, &paladin)
	var contendersDead = 0

	for len(contenders)-1 != contendersDead {
		contendersDead = battle(contenders, contendersDead)
		time.Sleep(time.Second * 3)
	}
	contenderWinner(contenders)
}
