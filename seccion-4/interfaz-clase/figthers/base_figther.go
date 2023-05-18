package figthers

type BaseFigther struct {
	Life int
}

func (bf BaseFigther) IsAlive() bool {
	return bf.Life > 0
}
