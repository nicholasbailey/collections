package collections

// A Stream is a lazy iterable that wraps an arbitrary
// iterator. Streams are rarely constructed directly,
// but are the conventional return type for most of the functional
// iteration methods (Map, Filter, etc.)
//
// Note that streams are generally *not* iterable multiple times, but
// provide no guards against multiple iteration. This may change in a
// future release where we may enforce single-time iteration
// per stream
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

func (stream *Stream) Any(matchFn func(interface{}) bool) bool {
	return anyHelper(stream, matchFn)
}
