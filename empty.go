package collections

func NewEmptyIterable() Iterable {
	return &EmptyIterable{}
}

// An Empty Iterable is an iterable that is guarenteed
// to have no elements. It's rarely used directly, but is
// a helpful tool for returning a more efficient 'empty'
// iterable for cases where we can guarentee an empty iterable
// (e.g. Skipping or Taking past the known length of a collection),
// mapping over an empty collection etc.
type EmptyIterable struct {
}

func (iterable *EmptyIterable) Iterator() Iterator {
	return &EmptyIterator{}
}

func (iterable *EmptyIterable) ForEach(iterFn func(interface{})) {

}

func (iterable *EmptyIterable) Map(mapFn func(interface{}) interface{}) Iterable {
	return iterable
}

func (iterable *EmptyIterable) Filter(filterFn func(interface{}) bool) Iterable {
	return iterable
}

func (iterable *EmptyIterable) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	return initialValue
}

func (iterable *EmptyIterable) ToSlice() []interface{} {
	return []interface{}{}
}

func (iterable *EmptyIterable) Take(count int) Iterable {
	return iterable
}

func (iterable *EmptyIterable) Skip(count int) Iterable {
	return iterable
}

func (iterable *EmptyIterable) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return iterable
}

func (iterable *EmptyIterable) Any(matchFn func(interface{}) bool) bool {
	return false
}

// An iterator with no elements
type EmptyIterator struct {
}

// Advances the Iterator by one step
func (empty *EmptyIterator) MoveNext() bool {
	return false
}

// Returns the current item in the Iterator
func (empty *EmptyIterator) Current() interface{} {
	panic(ErrIterationOutOfRange)
}
