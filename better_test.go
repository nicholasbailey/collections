package collections

import (
	"errors"
	"reflect"
	"testing"
)

// TODO - This should probably be replaced
// with Gomega or Assert

type expecter struct {
	t     *testing.T
	value interface{}
}

func expectFor(t *testing.T) func(interface{}) *expecter {
	t.Helper()
	return func(value interface{}) *expecter {
		return &expecter{
			value: value,
			t:     t,
		}
	}
}

func (expecter *expecter) ToBe(expected interface{}) {
	expecter.t.Helper()
	if expecter.value != expected {
		expecter.t.Fatalf("expected %v got %v", expected, expecter.value)
	}
}

func (expecter *expecter) ToDeepEqual(expected interface{}) {
	expecter.t.Helper()
	if !reflect.DeepEqual(expecter.value, expected) {
		expecter.t.Fatalf("expected %v got %v", expected, expecter.value)
	}
}

func (expecter *expecter) ToBeAssignableTo(expected reflect.Type) {
	expecter.t.Helper()
	if !reflect.TypeOf(expecter.value).AssignableTo(expected) {
		expecter.t.Fatalf("expected %v to be assignable to %v", expecter.value, expected)
	}
}

func (expecter *expecter) ToPanicWith(expected error) {
	expecter.t.Helper()
	f := expecter.value.(func())
	defer func() {
		expecter.t.Helper()
		err := recover()
		if err == nil {
			expecter.t.Fatalf("expect function to panic but it didn't")
		}
		if !errors.Is(err.(error), expected) {
			expecter.t.Fatalf("Failed with incorrect error. Expected %v, got %v", expected, err)
		}
	}()
	f()
}
