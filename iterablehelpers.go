package collections

// Helper functions to make it easier to implement the Iterable
// iterface with consistent logic

func forEachHelper(iterable Iterable, iterFn func(interface{})) {
	iterator := iterable.Iterator()
	for iterator.MoveNext() {
		iterFn(iterator.Current())
	}
}

func mapHelper(iterable Iterable, mapFn func(interface{}) interface{}) Iterable {
	oldIterator := iterable.Iterator()
	newIterator := &MapIterator{
		baseIterator: oldIterator,
		mapFn:        mapFn,
	}
	return NewStream(newIterator)
}

func filterHelper(iterable Iterable, filterFn func(interface{}) bool) Iterable {
	oldIterator := iterable.Iterator()
	newIterator := &FilterIterator{
		baseIterator: oldIterator,
		filterFn:     filterFn,
	}
	return NewStream(newIterator)
}

func foldHelper(iterable Iterable, initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	val := initialValue
	iterator := iterable.Iterator()
	for iterator.MoveNext() {
		val = reducerFn(val, iterator.Current())
	}
	return val
}

func toSliceHelper(iterable Iterable) []interface{} {
	slice := []interface{}{}
	iterator := iterable.Iterator()
	for iterator.MoveNext() {
		slice = append(slice, iterator.Current())
	}
	return slice
}

func takeHelper(interable Iterable, count int) Iterable {
	if count < 0 {
		panic(ErrInvalidTakeArgument)
	}
	iterator := &TakeIterator{
		baseIterator: interable.Iterator(),
		index:        0,
		count:        count,
	}

	return NewStream(iterator)
}

func skipHelper(iterable Iterable, count int) Iterable {
	if count < 0 {
		panic(ErrInvalidSkipArgument)
	}
	iterator := &SkipIterator{
		baseIterator: iterable.Iterator(),
		count:        count,
		consumed:     false,
	}
	return NewStream(iterator)
}

func skipWhileHelper(iterable Iterable, matchFn func(interface{}) bool) Iterable {
	iterator := &SkipWhileIterator{
		baseIterator: iterable.Iterator(),
		consumed:     false,
		matchFn:      matchFn,
	}
	return NewStream(iterator)
}

func anyHelper(iterable Iterable, matchFn func(interface{}) bool) bool {
	iterator := iterable.Iterator()
	for iterator.MoveNext() {
		if matchFn(iterator.Current()) {
			return true
		}
	}
	return false
}
