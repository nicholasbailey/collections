package collections

// An Iterable is a value that supports basic functional iteration
// operations. Iterables do not, in general, guarentee iteration order
// or that iteration will ever terminate. Iterables can also be destructive
// Meaning that iteration over the same iterable multipe times is not
// guarenteed to work.

type Iterable interface {
	// Iterates over the iterable and calls iterFn on each item
	// Note that if the iterable is infinite this will loop forever.
	ForEach(iterFn func(interface{}))

	// Returns an iterable whose items are the return value
	// of mapFn called on the items of the original iterable.
	Map(mapFn func(interface{}) interface{}) Iterable

	// Returns an iterable whose values are the items of the
	// original iterable for which filterFn returns true
	Filter(filterFn func(interface{}) bool) Iterable

	// Starting with initialValue calls reducerFn on
	// the current state of the fold and the next item of the iterable
	// and sets the state to the result. The end result is the
	// final return value of reducerFn.
	// Note that if the iterable is infinite, this will loop forever.
	Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{}

	// Gets an Iterator that iterates over the Iterable
	Iterator() Iterator

	// Returns a Slice whose elements are the items of the iterable.
	// Note that if the iterable is infinite, this will loop forever.
	ToSlice() []interface{}

	// // Returns an Iterable with the first "count" items of the Iterable
	// // If the iterable has fewer than "count" items, returns an iterable
	// // with all of them.
	// Take(count int) Iterable

	// // Returns an Iterable with the items of the Iterable except the first "count" items
	// // If the Iterable has fewer than "count" items, returns an empty iterable.
	// Skip(count int) Iterable

	// // Skips items in the iterable as long as skipFn is true, and retuns an
	// // Iterable with the remaining iterables.
	// SkipWhile(skipFn func(interface{}) bool)

	// GroupBy(groupFn func(interface{}) interface{}) Iterable

	// // Returns true if any items in the iterable. Note that
	// // if the iterable is infinite and matchFn is false for all items
	// // it will loop infinitely
	// Any(matchFn func(interface{}) bool) bool

	// // Returns the first item in the iterable matching the matchFn
	// // and boolean indicating if the value was found. Guarenteed
	// // to return the same result as Filter(matchFn).Head()
	// Find(matchFn func(interface{}) bool) (interface{}, bool)

	// // Returns an iterable with the elements of the original
	// // iterable sorted by the specified lessFn. The sort is not guarenteed
	// // to be stable. If the iterable is infinite, this will loop indefinitely.
	// SortBy(lessFn func(interface{}, interface{}) bool) Iterable

	// // Returns the first item in the iterable and
	// // and boolean indicating if collection was non-empty
	// Head() (interface{}, bool)

	// // Returns true if the iterable is empty, false otherwise. Note that
	// // you have to be very careful about calling this
	// // on destructive Iterables, since it may
	// // Iterate, rendering the Iterable unusable.
	// // Guarenteed to return the same results as Any(func(interface{}) bool {return true})
	// IsEmpty() bool

	// // Returns a vector whose elements are the items of the iterable.
	// // If the iterable is infinite, loops infinitely
	// ToVector() Vector

	// // Returns a Map whose keys are values returned by keyFn, and whose
	// // values are the values returned by valueFn. If multiple items return
	// // the same key, the last one will be used.
	// ToMap(keyFn func(interface{}) interface{}, valueFn func(interface{}) interface{}) Map

	// // Returns a map whose keys are values returned by keyFn, and whose
	// // values are the values returned by valueFn. If multiple items return
	// // the same key, the last one will be used.
	// ToGoMap(keyFn func(interface{}) interface{}, valueFn func(interface{}) interface{}) map[interface{}]interface{}
}

// An Iterator is a value that facilitates iteration logic.
// Iterators represent a moving pointer over
// some iterable.
//
// collections iterators use the following semantics.
// 1. The MoveNext method advances the iterator one step,
// and returns true if there are more items to iterate over
// and false if there are not
// 2. The Current method gets the current item
// 3. The iterator is assumed to be initialized in a
// 'pre-iteration state'. You must call 'MoveNext' once
// before you access Current. Calling Current() before
// calling MoveNext() should panic
// 4. If MoveNext() has returned false, Current() should
// panic
// 5. The correct way to loop over an iterator is.
//     for iter.MoveNext() {
//	       val := iter.Current()
//         ... Do stuff with value.
//     }
type Iterator interface {
	// Advances the iterator one step
	MoveNext() bool
	// Gets element currently pointed to by the Iterator
	Current() interface{}
}

// A FiniteIterable is an iterable with a fixed size.
type FiniteIterable interface {
	Iterable
	// The number of elements in the iterable
	Size() int
}

// A Sequence is an immutable iterable with a fixed length and order.
// Examples of Sequence types include LinkedLists, SliceSequences
// and Vectors. All Sequences are zero-indexed
type Sequence interface {
	// All Sequences are also FiniteIterables
	FiniteIterable

	// Gets the element at the specified index. Panics with
	// ErrIndexOutOfRange if the index is out of range.
	Get(index int) interface{}

	// Functional update.
	// Creates a copy of the sequence that is identical
	// with the specified value changed at the specified
	// index
	Update(index int, value interface{}) Sequence

	// Functional Append.
	// Creates a copy of the sequence with a new value at the
	// end of the collection
	Append(value interface{}) Sequence

	// Functional Prepend.
	// Creates a copy of the sequence with a new value at the
	// start of the collection
	Prepend(value interface{}) Sequence

	// Slice. Creates a new sequence
	Slice(start int, end int) Sequence
}

// A Set is an immutable iterable with no duplicates. Sets support standard
// set operations. Sets do not in general guarentee iteration
// order. Implementations that do will also support 'Sequence' semantics
type Set interface {
	FiniteIterable

	Contains(value interface{}) bool
	SubsetOf(other Set) bool
	Add(value interface{}) Set
	Remove(value interface{}) Set
	Intersect(other Set) Set
	Union(other Set) Set
	Difference(other Set) Set
}

// A Map is an immutable iterable of Key-Value pairs supporting
// Key-Value store semantics. Maps do not in general guarentee
// iteration order, though many implementations do
type Map interface {
	Iterable
	Contains(key interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{}) Map
	Remove(key interface{}) Map
	Merge(other Map) Map
	Keys() Iterable
	Values() Iterable
	KeySet() Set
}
