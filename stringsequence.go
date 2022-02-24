package collections

// Returns a sequence which provides rune-by-rune
// iteration over the string. Currently actually returns
// a slice sequence of a rune slice of the string.
// Since this does mean copying the data of the string,
// this may not be the best option for extremely large strings
// or performance critical code.
func NewStringSequence(str string) Sequence {
	slice := []interface{}{}

	for _, char := range str {
		slice = append(slice, char)
	}
	return NewSliceSequence(slice...)
}
