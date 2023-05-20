package figthers

import "math/rand"

type Criminal struct {
	BaseFigther
}

func (c Criminal) ThrowAttack() int {
	return rand.Intn(13)
}

func (c *Criminal) RecieveAttack(intensity int) {
	c.Life -= intensity
}

func (c Criminal) GetName() string {
	return "Criminal"
}

//method  signature = nombre  + parametros + valor de retorno
//method set  es el conjunto de metodos
