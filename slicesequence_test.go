package collections

import (
	"strings"
	"testing"
)

func TestEmptySliceSequenceSize(t *testing.T) {
	expect := expectFor(t)
	sequence := NewSliceSequence()
	expect(sequence.Size()).ToBe(0)
}

func TestSliceSequenceAppendIsImmutable(t *testing.T) {
	expect := expectFor(t)

	sequences := []Sequence{}
	var seq Sequence = NewSliceSequence()
	// Build a slice of sequences, each of which has
	// n elements, where n is the index in the slice
	for i := 0; i < 1000; i++ {
		sequences = append(sequences, seq)
		seq = seq.Append(i)
	}

	for i := 0; i < 1000; i++ {
		seq := sequences[i]
		expect(seq.Size()).ToBe(i)
		for j := 0; j < i; j++ {
			expect(seq.Get(j)).ToBe(j)
		}
	}
}

func TestSliceSequnceIndexOutOfRange(t *testing.T) {
	expect := expectFor(t)
	slice := []interface{}{1, 2, 3}
	seq := NewSliceSequence(slice...)

	expect(func() { seq.Get(4) }).ToPanicWith(ErrIndexOutOfRange)
	expect(func() { seq.Get(-1) }).ToPanicWith(ErrIndexOutOfRange)
	expect(func() { seq.Update(-1, -1) }).ToPanicWith(ErrIndexOutOfRange)
	expect(func() { seq.Update(-1, 4) }).ToPanicWith(ErrIndexOutOfRange)
	expect(seq.ToSlice()).ToDeepEqual([]interface{}{1, 2, 3})
}

func TestSliceSequenceAppendDoesNotMutateUnderlyingSlice(t *testing.T) {
	expect := expectFor(t)
	slice := []interface{}{"Picture", "yourself", "on", "a", "boat", "in", "a", "river"}
	original := NewSliceSequence(slice...)
	_ = original.Append("with")
	expect(slice).ToDeepEqual([]interface{}{"Picture", "yourself", "on", "a", "boat", "in", "a", "river"})
}

func TestSliceSequenceGet(t *testing.T) {
	expect := expectFor(t)
	slice := []interface{}{"Picture", "yourself", "on", "a", "boat", "in", "a", "river"}
	seq := NewSliceSequence(slice...)

	expect(seq.Get(1)).ToBe("yourself")
}

func TestSliceSequenceUpdate(t *testing.T) {
	expect := expectFor(t)

	slice := []interface{}{"Picture", "yourself", "on", "a", "boat", "in", "a", "river"}
	original := NewSliceSequence(slice...)

	new := original.Update(1, 4.5)
	expect(new.Size()).ToBe(8)
	expect(new.Get(1)).ToBe(4.5)
	expect(original.Get(1)).ToBe("yourself")
	expect(slice).ToDeepEqual([]interface{}{"Picture", "yourself", "on", "a", "boat", "in", "a", "river"})
}

func TestSliceSequenceMap(t *testing.T) {
	expect := expectFor(t)

	seq := NewSliceSequence(1, 2, 3, 4, 5, 6)

	actual := seq.Map(func(x interface{}) interface{} { return x.(int) * x.(int) }).ToSlice()

	expect(actual).ToDeepEqual([]interface{}{1, 4, 9, 16, 25, 36})
}

func TestSliceSequenceFilter(t *testing.T) {
	expect := expectFor(t)

	seq := NewSliceSequence(1, 2, 3, 4, 5, 6)

	actual := seq.Filter(func(x interface{}) bool { return x.(int)%2 == 0 }).ToSlice()

	expect(actual).ToDeepEqual([]interface{}{2, 4, 6})
}

func TestSliceSequenceForEach(t *testing.T) {
	expect := expectFor(t)

	seq := NewSliceSequence("Too", "hot!", "Hot", "Damn!")

	result := []string{}
	fn := func(v interface{}) {
		if strings.Contains(v.(string), "!") {
			result = append(result, v.(string))
		}
	}

	seq.ForEach(fn)

	expect(result).ToDeepEqual([]string{"hot!", "Damn!"})
}

func TestSliceSequenceFold(t *testing.T) {
	expect := expectFor(t)
	seq := NewSliceSequence("Too", "hot!", "Hot", "Damn!")

	result := seq.Fold("", func(x interface{}, y interface{}) interface{} {
		if x == "" {
			return y
		} else {
			return x.(string) + " " + y.(string)
		}
	})

	expect(result).ToBe("Too hot! Hot Damn!")
}

func TestSliceSequenceAny(t *testing.T) {
	expect := expectFor(t)
	seq1 := NewSliceSequence(float32(1), 1.5, int8(1), 'a')
	seq2 := NewSliceSequence(float32(1), 1.5, int8(1))
	fn := func(v interface{}) bool {
		switch v.(type) {
		case rune:
			return true
		default:
			return false
		}
	}
	expect(seq1.Any(fn)).ToBe(true)
	expect(seq2.Any(fn)).ToBe(false)
}

func TestSliceSequenceSkip(t *testing.T) {
	expect := expectFor(t)
	seq := NewSliceSequence("some", "words", "that", "mean", "something")
	expect(seq.Skip(3).ToSlice()).ToDeepEqual([]interface{}{"mean", "something"})
	expect(seq.Skip(8).ToSlice()).ToDeepEqual([]interface{}{})
	expect(func() { seq.Skip(-1) }).ToPanicWith(ErrInvalidSkipArgument)
}

func TestSliceSequenceTake(t *testing.T) {
	expect := expectFor(t)
	seq := NewSliceSequence("Call", "me", "Ishmael", "Some", "years", "ago")

	expect(seq.Take(3).ToSlice()).ToDeepEqual([]interface{}{"Call", "me", "Ishmael"})
	expect(seq.Take(30).ToSlice()).ToDeepEqual([]interface{}{"Call", "me", "Ishmael", "Some", "years", "ago"})
	expect(func() { seq.Take(-90) }).ToPanicWith(ErrInvalidTakeArgument)
}

func TestSliceSequenceSkipWhile(t *testing.T) {
	expect := expectFor(t)

	seq := NewSliceSequence("Call", "me", "Ishmael", "Some", "years", "ago")

	matchFn := func(v interface{}) bool {
		return len(v.(string)) < 5
	}

	expect(seq.SkipWhile(matchFn).ToSlice()).ToDeepEqual([]interface{}{"Ishmael", "Some", "years", "ago"})
}
