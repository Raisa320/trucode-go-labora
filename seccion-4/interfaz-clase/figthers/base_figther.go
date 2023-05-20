package figthers

type BaseFigther struct {
	Life        int
	LifeInitial int
}

func (bf BaseFigther) IsAlive() bool {
	return bf.Life > 0
}

func (bf *BaseFigther) ActualLife() int {
	return bf.Life
}
