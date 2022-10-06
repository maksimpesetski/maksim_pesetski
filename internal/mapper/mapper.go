package mapper

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// MapString accepts input that implements Interface and calls i.TransformRune in each rune of GetValueAsRuneSlice().
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}
