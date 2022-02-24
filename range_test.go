package collections

import "testing"

func TestRangeToSlice(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(2, 11)
	s := rng.ToSlice()
	expect(s).ToDeepEqual([]interface{}{2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestRangeCanHaveNegativeBounds(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(-20, -10)
	s := rng.ToSlice()
	expect(s).ToDeepEqual([]interface{}{-20, -19, -18, -17, -16, -15, -14, -13, -12, -11})
}

func TestRangeMap(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(2, 6)
	mapped := rng.Map(func(v interface{}) interface{} { return v.(int) * v.(int) }).ToSlice()

	expect(mapped).ToDeepEqual([]interface{}{4, 9, 16, 25})
}

func TestRangeFilter(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 100)
	filtered := rng.Filter(func(v interface{}) bool { return v.(int)%5 == 0 }).ToSlice()

	expect(filtered).ToDeepEqual([]interface{}{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95})
}

func TestRangeForEach(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(1, 8)
	actual := []interface{}{}
	rng.ForEach(func(val interface{}) {
		actual = append(actual, val)
	})
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7}
	expect(expected).ToDeepEqual(actual)
}

func TestRangeFold(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 10)
	actual := rng.Fold(4, func(exist interface{}, next interface{}) interface{} {
		return exist.(int) + next.(int)
	})
	expect(actual).ToBe(4 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9)
}

func TestRangeNewInvalidBounds(t *testing.T) {
	expect := expectFor(t)
	expect(func() { NewRange(10, 2) }).ToPanicWith(ErrInvalidRangeBounds)
}

func TestRangeAny(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(-10, -1)
	expect(rng.Any(func(v interface{}) bool { return v.(int) > 0 })).ToBe(false)
}

func TestRangeSkip(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 20)
	expected := []interface{}{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	expect(rng.Skip(10).ToSlice()).ToDeepEqual(expected)
}

func TestRangeSkipMoreThanElemnts(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 20)
	expect(rng.Skip(21).ToSlice()).ToDeepEqual([]interface{}{})
}

func TestRangeTake(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 1000)
	expected := []interface{}{0, 1, 2, 3}
	expect(rng.Take(4).ToSlice()).ToDeepEqual(expected)
}

func TestRangeTakeLongerThanList(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 5)
	expected := []interface{}{0, 1, 2, 3, 4}
	expect(rng.Take(30).ToSlice()).ToDeepEqual(expected)
}

func TestRangeSkipWhile(t *testing.T) {
	expect := expectFor(t)
	rng := NewRange(0, 10)
	expected := []interface{}{5, 6, 7, 8, 9}
	actual := rng.SkipWhile(func(v interface{}) bool { return v.(int) < 5 }).ToSlice()
	expect(actual).ToDeepEqual(expected)
}
