package collections

// This file contains assorted iterators. Most are adapter iterators
// which wrap a base iterator and modify its behavior.
// This is the standard way of supporting functions like Map and Filter

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

// An Iterator backed by a sequence, which handles iteration
// through calls to Sequence.Get. Note that this is not
// guarenteed to be the most efficient way to iterate through
// the Sequence, and many Sequence types implement their own
// custom Iterators
type SequenceIterator struct {
	index    int
	sequence Sequence
}

// Creates a new SequenceIterator from a sequence
func NewSequenceIterator(sequence Sequence) *SequenceIterator {
	return &SequenceIterator{
		sequence: sequence,
		index:    -1,
	}
}

// Advances the Iterator by one step
func (iterator *SequenceIterator) MoveNext() bool {
	iterator.index += 1
	return iterator.index < iterator.sequence.Size()
}

// Returns the current item from the iterator
func (iterator *SequenceIterator) Current() interface{} {
	if iterator.index < iterator.sequence.Size() {
		panic(ErrIterationOutOfRange)
	}
	return iterator.sequence.Get(iterator.index)
}

// An iterator that lazily evaluates a Map operation
// on a base iterator
type MapIterator struct {
	baseIterator Iterator
	mapFn        func(interface{}) interface{}
}

// Advances the Iterator by one step
func (iterator *MapIterator) MoveNext() bool {
	return iterator.baseIterator.MoveNext()
}

// Returns the current item from the iterator
func (iterator *MapIterator) Current() interface{} {
	return iterator.mapFn(iterator.baseIterator.Current())
}

// An iteratore which lazily evaluates a Filter operation
// on a base iterator
type FilterIterator struct {
	baseIterator Iterator
	filterFn     func(interface{}) bool
}

// Advances the Iterator by one step
func (iterator *FilterIterator) MoveNext() bool {
	base := iterator.baseIterator
	for base.MoveNext() {
		current := base.Current()
		if iterator.filterFn(current) {
			return true
		}
	}
	return false
}

// Returns the current item from the iterator
func (iterator *FilterIterator) Current() interface{} {
	return iterator.baseIterator.Current()
}

// An iterator which lazily evaluates a Take operation
// on a base iterator
type TakeIterator struct {
	baseIterator Iterator
	count        int
	index        int
}

func (iterator *TakeIterator) MoveNext() bool {
	if iterator.index >= iterator.count {
		return false
	}
	moveNext := iterator.baseIterator.MoveNext()
	if moveNext {
		iterator.index += 1
	}
	return moveNext
}

func (iterator *TakeIterator) Current() interface{} {
	return iterator.baseIterator.Current()
}

// An iterator which lazily evaluates a Skip
// operation on base iterator
type SkipIterator struct {
	baseIterator Iterator
	consumed     bool
	count        int
}

func (iterator *SkipIterator) MoveNext() bool {
	if !iterator.consumed {
		iterator.consumed = true
		hasNext := false
		for i := 0; i <= iterator.count; i++ {
			hasNext = iterator.baseIterator.MoveNext()
		}
		return hasNext
	}
	return iterator.baseIterator.MoveNext()
}

func (iterator *SkipIterator) Current() interface{} {
	return iterator.baseIterator.Current()
}

// An Iterator that lazily evaluates a
// SkipWhile operation on a base iterator
type SkipWhileIterator struct {
	baseIterator Iterator
	consumed     bool
	matchFn      func(interface{}) bool
}

func (iterator *SkipWhileIterator) MoveNext() bool {
	if !iterator.consumed {
		iterator.consumed = true
		for iterator.baseIterator.MoveNext() {
			current := iterator.baseIterator.Current()
			if !iterator.matchFn(current) {
				return true
			}
		}
		return false
	}

	return iterator.baseIterator.MoveNext()
}

func (iterator *SkipWhileIterator) Current() interface{} {
	return iterator.baseIterator.Current()
}
