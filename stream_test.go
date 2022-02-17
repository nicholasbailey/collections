package collections

import (
	"reflect"
	"strings"
	"testing"
)

func buildStream(data []interface{}) *Stream {
	return NewStream(NewSliceIterator(data))
}

func TestMapSimple(t *testing.T) {
	data := []interface{}{"A", "B", "C", "D", "E"}
	stream := buildStream(data)

	mapFn := func(v interface{}) interface{} {
		return strings.ToLower(v.(string))
	}

	expected := []interface{}{"a", "b", "c", "d", "e"}
	actual := stream.Map(mapFn).ToSlice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}

func TestFilterSimple(t *testing.T) {
	data := []interface{}{5, 2, 3, 4, 4}
	stream := buildStream(data)
	filterFn := func(v interface{}) bool {
		return v.(int) > 3
	}
	expected := []interface{}{5, 4, 4}
	actual := stream.Filter(filterFn).ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}

func TestFoldSimple(t *testing.T) {
	data := []interface{}{"a", "b", "c", "d", "e"}
	stream := buildStream(data)
	reducerFn := func(state interface{}, next interface{}) interface{} {
		return state.(string) + next.(string)
	}
	expected := "1abcde"
	actual := stream.Fold("1", reducerFn)
	if actual != expected {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}

func TestForEachSimple(t *testing.T) {
	data := []interface{}{"a", "b", "c", "d", "e"}
	stream := buildStream(data)

	actual := ""

	stream.ForEach(func(x interface{}) {
		actual = actual + x.(string)
	})

	expected := "abcde"

	if actual != expected {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}

func TestTakeSimple(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	stream := buildStream(data)

	actual := stream.Take(5).ToSlice()
	expected := []interface{}{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestSkipSimple(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	stream := buildStream(data)

	actual := stream.Skip(3).ToSlice()
	expected := []interface{}{4, 5, 6, 7, 8, 9, 10, 11}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestSkipWhileSimple(t *testing.T) {
	data := []interface{}{2, 2, 2, 4, 2, 6, 7, 2, 9, 10, 11}
	stream := buildStream(data)
	actual := stream.SkipWhile(func(v interface{}) bool { return v.(int) == 2 }).ToSlice()
	expected := []interface{}{4, 2, 6, 7, 2, 9, 10, 11}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
