package figthers

//interfaz conunto de metodos,  firma, parametros y que devuelve
type Contender interface {
	ThrowAttack() int
	RecieveAttack(intensity int)
	IsAlive() bool
	GetName() string
}
