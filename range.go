package collections

// A Range is an iterable which yields successive integer values
// between two endpoints. Ranges are non-destructive and can be
// safely iterated repeatedly
type Range struct {
	begin int
	end   int
}

func NewRange(begin int, end int) *Range {
	if end < begin {
		panic(ErrInvalidRangeBounds)
	}
	return &Range{
		begin: begin,
		end:   end,
	}
}

func (rng *Range) Iterator() Iterator {
	return &RangeIterator{
		current: rng.begin - 1,
		end:     rng.end,
	}
}

type RangeIterator struct {
	end     int
	current int
}

func (rangeIterator *RangeIterator) MoveNext() bool {
	rangeIterator.current += 1
	return rangeIterator.current < rangeIterator.end
}

func (rangeIterator *RangeIterator) Current() interface{} {
	return rangeIterator.current
}

func (rng *Range) ToSlice() []interface{} {
	slice := make([]interface{}, rng.end-rng.begin)
	for val := rng.begin; val < rng.end; val++ {
		index := val - rng.begin
		slice[index] = val
	}
	return slice
}

func (rng *Range) Map(mapFn func(interface{}) interface{}) Iterable {
	return mapHelper(rng, mapFn)
}

func (rng *Range) Filter(filterFn func(interface{}) bool) Iterable {
	return filterHelper(rng, filterFn)
}

func (rng *Range) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	val := initialValue
	for i := rng.begin; i < rng.end; i++ {
		val = reducerFn(val, i)
	}
	return val
}

func (rng *Range) Take(count int) Iterable {
	length := rng.end - rng.begin
	if count >= length {
		return rng
	}
	return NewRange(rng.begin, rng.begin+count)
}

func (rng *Range) Skip(count int) Iterable {
	length := rng.end - rng.begin
	if count >= length {
		return &EmptyIterable{}
	}
	return NewRange(rng.begin+count, rng.end)
}

func (rng *Range) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return skipWhileHelper(rng, matchFn)
}

func (rng *Range) Any(matchFn func(interface{}) bool) bool {
	for i := rng.begin; i < rng.end; i++ {
		if matchFn(i) {
			return true
		}
	}
	return false
}

func (rng *Range) ForEach(iterFn func(interface{})) {
	for i := rng.begin; i < rng.end; i++ {
		iterFn(i)
	}
}
