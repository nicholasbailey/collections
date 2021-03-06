package collections

import (
	"strings"
	"testing"
)

func buildStream(data []interface{}) *Stream {
	return NewStream(NewSliceIterator(data))
}

func TestStreamMapSimple(t *testing.T) {
	expect := expectFor(t)
	data := []interface{}{"A", "B", "C", "D", "E"}
	stream := buildStream(data)

	mapFn := func(v interface{}) interface{} {
		return strings.ToLower(v.(string))
	}

	expected := []interface{}{"a", "b", "c", "d", "e"}
	actual := stream.Map(mapFn).ToSlice()

	expect(actual).ToDeepEqual(expected)
}

func TestStreamMapSwapTypes(t *testing.T) {
	expect := expectFor(t)
	data := []interface{}{"Apple", "Bear", "Cot", "Dinosaur", "Em"}
	stream := buildStream(data)

	mapFn := func(v interface{}) interface{} {
		return len(v.(string))
	}

	expected := []interface{}{5, 4, 3, 8, 2}
	actual := stream.Map(mapFn).ToSlice()

	expect(actual).ToDeepEqual(expected)
}

func TestStreamFilterSimple(t *testing.T) {
	expect := expectFor(t)

	data := []interface{}{5, 2, 3, 4, 4}
	stream := buildStream(data)
	filterFn := func(v interface{}) bool {
		return v.(int) > 3
	}
	expected := []interface{}{5, 4, 4}
	actual := stream.Filter(filterFn).ToSlice()

	expect(actual).ToDeepEqual(expected)
}

func TestStreamFilterNoMatches(t *testing.T) {
	expect := expectFor(t)

	data := []interface{}{5, 2, 3, 4, 4}
	stream := buildStream(data)
	filterFn := func(v interface{}) bool {
		return false
	}
	expected := []interface{}{}
	actual := stream.Filter(filterFn).ToSlice()

	expect(actual).ToDeepEqual(expected)
}

func TestStreamFoldSimple(t *testing.T) {
	expect := expectFor(t)

	data := []interface{}{"a", "b", "c", "d", "e"}
	stream := buildStream(data)
	reducerFn := func(state interface{}, next interface{}) interface{} {
		return state.(string) + next.(string)
	}
	expected := "1abcde"
	actual := stream.Fold("1", reducerFn)

	expect(actual).ToBe(expected)
}

func TestStreamForEachSimple(t *testing.T) {
	expect := expectFor(t)

	data := []interface{}{"a", "b", "c", "d", "e"}
	stream := buildStream(data)

	actual := ""

	stream.ForEach(func(x interface{}) {
		actual = actual + x.(string)
	})

	expected := "abcde"

	expect(actual).ToBe(expected)
}

func TestStreamTakeSimple(t *testing.T) {
	expect := expectFor(t)

	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	stream := buildStream(data)

	actual := stream.Take(5).ToSlice()
	expected := []interface{}{1, 2, 3, 4, 5}

	expect(actual).ToDeepEqual(expected)
	expect(func() { buildStream(data).Take(-5) }).ToPanicWith(ErrInvalidTakeArgument)
}

func TestStreamSkipSimple(t *testing.T) {
	expect := expectFor(t)
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	stream := buildStream(data)

	actual := stream.Skip(3).ToSlice()
	expected := []interface{}{4, 5, 6, 7, 8, 9, 10, 11}

	expect(actual).ToDeepEqual(expected)
	expect(func() { buildStream(data).Skip(-5) }).ToPanicWith(ErrInvalidSkipArgument)
}

func TestStreamSkipWhileSimple(t *testing.T) {
	expect := expectFor(t)
	data := []interface{}{2, 2, 2, 4, 2, 6, 7, 2, 9, 10, 11}
	stream := buildStream(data)
	actual := stream.SkipWhile(func(v interface{}) bool { return v.(int) == 2 }).ToSlice()
	expected := []interface{}{4, 2, 6, 7, 2, 9, 10, 11}

	expect(actual).ToDeepEqual(expected)
}
