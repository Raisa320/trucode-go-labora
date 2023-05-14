package slices

func RotateRightMyVersion(cadena string) string {
	if len(cadena) > 1 {
		letraRotar := string(cadena[len(cadena)-1])
		cadena = cadena[0 : len(cadena)-1]
		rotado := append([]byte(letraRotar), []byte(cadena)...)
		return string(rotado)
	}
	return cadena
}

func RotateLeftMyVersion(cadena string) string {
	if len(cadena) > 1 {
		letraRotar := string(cadena[0])
		cadena = cadena[1:]
		rotado := append([]byte(cadena), []byte(letraRotar)...)
		return string(rotado)
	}
	return cadena
}

func RotateRightVituVersion(text string) string { // graphical idea: https://jamboard.google.com/d/1h5aAO-wiPI3TO5QIChnGQEaAPlT8dU79k3ZFckh2NW8/viewer?mtt=v7wp3cnbz3an&f=0
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}

func RotateLeftVituVersion(text string) string {
	size := len(text)
	rotatedText := ""
	for i := range text {
		char := string(text[(size+1+i)%size])
		rotatedText = rotatedText + char
	}
	return rotatedText
}
