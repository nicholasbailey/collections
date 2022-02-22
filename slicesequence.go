package collections

// An Immutable Sequence which wraps a slice. The main reason
// for using a SliceSequence is to support Iterable behavior on
// slices.
// Access and iterations are highly efficient on SliceSequences because
// they reduce to slice indexing.
// Immutable data operations on SliceSequences are inefficient
// (O(n) for Update, Append, Prepend) because SliceSequences do not
// support data reuse.
type SliceSequence struct {
	slice []interface{}
}

// Factory for SliceSequences
func NewSliceSequence(values ...interface{}) *SliceSequence {
	return &SliceSequence{
		slice: values,
	}
}

// Iterator Methods

func (sliceSequence *SliceSequence) Iterator() Iterator {
	return NewSliceIterator(sliceSequence.slice)
}

func (sliceSequence *SliceSequence) ForEach(iterFn func(interface{})) {
	forEachHelper(sliceSequence, iterFn)
}

func (sliceSequence *SliceSequence) Map(mapFn func(interface{}) interface{}) Iterable {
	return mapHelper(sliceSequence, mapFn)
}

func (sliceSequence *SliceSequence) Filter(filterFn func(interface{}) bool) Iterable {
	return filterHelper(sliceSequence, filterFn)
}

func (sliceSequence *SliceSequence) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	return foldHelper(sliceSequence, initialValue, reducerFn)
}

func (sliceSequence *SliceSequence) ToSlice() []interface{} {
	return toSliceHelper(sliceSequence)
}

func (sliceSequence *SliceSequence) Take(count int) Iterable {
	return takeHelper(sliceSequence, count)
}

func (sliceSequence *SliceSequence) Skip(count int) Iterable {
	return skipHelper(sliceSequence, count)
}

func (sliceSequence *SliceSequence) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return skipWhileHelper(sliceSequence, matchFn)
}

func (sliceSequence *SliceSequence) Any(matchFn func(interface{}) bool) bool {
	return anyHelper(sliceSequence, matchFn)
}

// Sequence Methods
func (sliceSequence *SliceSequence) Size() int {
	return len(sliceSequence.slice)
}

func (sliceSequence *SliceSequence) Append(value interface{}) Sequence {
	originalLength := len(sliceSequence.slice)
	newSlice := make([]interface{}, originalLength+1)
	copy(newSlice, sliceSequence.slice)
	newSlice[originalLength] = value
	return NewSliceSequence(newSlice...)
}

func (sliceSequence *SliceSequence) Get(index int) interface{} {
	if index < 0 || index >= sliceSequence.Size() {
		panic(ErrIndexOutOfRange)
	}
	return sliceSequence.slice[index]
}

func (sliceSequence *SliceSequence) Update(index int, value interface{}) Sequence {
	if index < 0 || index >= sliceSequence.Size() {
		panic(ErrIndexOutOfRange)
	}
	originalLength := len(sliceSequence.slice)
	newSlice := make([]interface{}, originalLength)
	copy(newSlice, sliceSequence.slice)
	newSlice[index] = value
	return NewSliceSequence(newSlice...)
}

type SliceIterator struct {
	slice []interface{}
	index int
}

func NewSliceIterator(slice []interface{}) *SliceIterator {
	return &SliceIterator{
		slice: slice,
		index: -1,
	}
}

func (iterator *SliceIterator) MoveNext() bool {
	iterator.index += 1
	return iterator.index < len(iterator.slice)
}

func (iterator *SliceIterator) Current() interface{} {
	if iterator.index >= len(iterator.slice) {
		panic(ErrIterationOutOfRange)
	}
	return iterator.slice[iterator.index]
}
