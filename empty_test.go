package collections

import "testing"

func TestEmptyIterableIterator(t *testing.T) {
	iterable := NewEmptyIterable()
	iterator := iterable.Iterator()
	for iterator.MoveNext() {
		t.Fatalf("An empty iterator should never iterate")
	}
}

func TestEmptyIterableForEach(t *testing.T) {
	iterable := NewEmptyIterable()
	iterable.ForEach(func(interface{}) {
		t.Fatalf("An empty iteratable should never iterate")
	})
}

func TestEmptyIterableMap(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	actual := iterable.Map(func(v interface{}) interface{} {
		t.Fatalf("An empty iterable should never map")
		return v
	}).ToSlice()
	expect(actual).ToDeepEqual([]interface{}{})
}

func TestEmptyIterableFilter(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	actual := iterable.Filter(func(v interface{}) bool {
		t.Fatalf("An empty iterable should never map")
		return true
	}).ToSlice()
	expect(actual).ToDeepEqual([]interface{}{})
}

func TestEmptyIterableToSlice(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	expect(iterable.ToSlice()).ToDeepEqual([]interface{}{})
}

func TestEmptyIterableTake(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	expect(iterable.Take(5).ToSlice()).ToDeepEqual([]interface{}{})
}

func TestEmptyIterableSkip(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	expect(iterable.Skip(5).ToSlice()).ToDeepEqual([]interface{}{})
}

func TestEmptyIterableFold(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	val := iterable.Fold("base", func(x interface{}, y interface{}) interface{} {
		t.Fatal("An empty iterable should never fold")
		return y
	})

	expect(val).ToBe("base")
}

func TestEmptyIterableAny(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	result := iterable.Any(func(v interface{}) bool {
		t.Fatal("An empty iterable should never any")
		return true
	})
	expect(result).ToBe(false)
}

func TestEmptyIterableSkipWhile(t *testing.T) {
	expect := expectFor(t)
	iterable := NewEmptyIterable()
	result := iterable.SkipWhile(func(v interface{}) bool {
		t.Fatal("An empty iterable should never any")
		return true
	}).ToSlice()
	expect(result).ToDeepEqual([]interface{}{})
}
