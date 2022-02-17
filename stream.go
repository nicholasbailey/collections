package collections

type Stream struct {
	iterator Iterator
}

func NewStream(iterator Iterator) *Stream {
	return &Stream{
		iterator: iterator,
	}
}

func (iterable *Stream) Iterator() Iterator {
	return iterable.iterator
}

func (iterable *Stream) ForEach(iterFn func(interface{})) {
	forEachHelper(iterable, iterFn)
}

func (iterable *Stream) Map(mapFn func(interface{}) interface{}) Iterable {
	return mapHelper(iterable, mapFn)
}

func (iterable *Stream) Filter(filterFn func(interface{}) bool) Iterable {
	return filterHelper(iterable, filterFn)
}

func (iterable *Stream) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	return foldHelper(iterable, initialValue, reducerFn)
}

func (iterable *Stream) ToSlice() []interface{} {
	return toSliceHelper(iterable)
}

func (iterable *Stream) Take(count int) Iterable {
	return takeHelper(iterable, count)
}

func (stream *Stream) Skip(count int) Iterable {
	return skipHelper(stream, count)
}

func (stream *Stream) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return skipWhileHelper(stream, matchFn)
}
