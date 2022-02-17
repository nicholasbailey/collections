package collections

type EmptyIterator struct {
}

func (empty *EmptyIterator) MoveNext() bool {
	return false
}

func (empty *EmptyIterator) Current() interface{} {
	panic(ErrIterationOutOfRange)
}

type SequenceIterator struct {
	index    int
	sequence Sequence
}

func NewSequenceIterator(sequence Sequence) Iterator {
	return &SequenceIterator{
		sequence: sequence,
		index:    -1,
	}
}

func (iterator *SequenceIterator) MoveNext() bool {
	iterator.index += 1
	return iterator.index < iterator.sequence.Size()
}

func (iterator *SequenceIterator) Current() interface{} {
	if iterator.index < iterator.sequence.Size() {
		panic(ErrIterationOutOfRange)
	}
	return iterator.sequence.Get(iterator.index)
}

type MapIterator struct {
	baseIterator Iterator
	mapFn        func(interface{}) interface{}
}

func (iterator *MapIterator) MoveNext() bool {
	return iterator.baseIterator.MoveNext()
}

func (iterator *MapIterator) Current() interface{} {
	return iterator.mapFn(iterator.baseIterator.Current())
}

type FilterIterator struct {
	baseIterator Iterator
	filterFn     func(interface{}) bool
}

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

func (iterator *FilterIterator) Current() interface{} {
	return iterator.baseIterator.Current()
}

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
