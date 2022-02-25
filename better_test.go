package collections

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// TODO - This should probably be replaced
// with Gomega or Assert

type expecter struct {
	t      *testing.T
	value  interface{}
	invert bool
}

func expectFor(t *testing.T) func(interface{}) *expecter {
	t.Helper()
	return func(value interface{}) *expecter {
		return &expecter{
			value:  value,
			t:      t,
			invert: false,
		}
	}
}

func (expect *expecter) Not() *expecter {
	expect.t.Helper()
	return &expecter{
		value:  expect.value,
		t:      expect.t,
		invert: true,
	}
}

func (expect *expecter) doCheck(checkFn func() bool, msgFn func() string) {
	expect.t.Helper()

	var passed bool
	if expect.invert {
		passed = !checkFn()
	} else {
		passed = checkFn()
	}
	if !passed {
		var prefix string
		if expect.invert {
			prefix = "did not expect"
		} else {
			prefix = "expected"
		}
		errMsg := prefix + " " + msgFn()
		expect.t.Fatal(errMsg)
	}
}

func (expecter *expecter) ToBe(expected interface{}) {
	expecter.t.Helper()
	expecter.doCheck(
		func() bool {
			return expecter.value == expected
		},
		func() string { return fmt.Sprintf("%v to equal %v", expected, expecter.value) },
	)
}

func (expecter *expecter) ToDeepEqual(expected interface{}) {
	expecter.t.Helper()
	expecter.doCheck(
		func() bool { return reflect.DeepEqual(expecter.value, expected) },
		func() string { return fmt.Sprintf("%v to deep equal %v", expected, expecter.value) },
	)
}

func (expecter *expecter) ToBeAssignableTo(expected reflect.Type) {
	expecter.t.Helper()
	expecter.doCheck(
		func() bool { return reflect.TypeOf(expecter.value).AssignableTo(expected) },
		func() string { return fmt.Sprintf("%v to be assignable to %v", expected, expecter.value) },
	)
}

func (expecter *expecter) ToPanicWith(expected error) {
	expecter.t.Helper()
	if !expecter.invert {
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
	} else {
		f := expecter.value.(func())
		defer func() {
			expecter.t.Helper()
			err := recover()
			if err != nil {
				expecter.t.Fatalf("expect function not panic but it did")
			}
		}()
		f()
	}

}
