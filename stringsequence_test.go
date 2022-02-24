package collections

import (
	"reflect"
	"testing"
)

func TestStringSequenceIsActuallyASliceSequence(t *testing.T) {
	expect := expectFor(t)
	seq := NewStringSequence("Hello")

	expect(seq.Size()).ToBe(5)
	expect(reflect.TypeOf(seq).AssignableTo(reflect.TypeOf(NewSliceSequence()))).ToBe(true)
	expect(seq.ToSlice()).ToDeepEqual([]interface{}{'H', 'e', 'l', 'l', 'o'})
}
