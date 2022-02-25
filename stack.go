package collections

import "fmt"

// Note: We very deliberately didn't call this type a List
// because we don't want users of the library making the
// mistake of using it as a general purpose sequence.
// The right type for a general purpose sequence is Vector

// A immutable stack implemented as a singly linked list
// Stacks implement Sequence, as well as an additional
// Stack interface with the top of the
// stack being the first element.
type Stack interface {
	Sequence

	// Returns a new Stack with the new item
	// pushed to the top
	Push(item interface{}) Stack

	// Returns a new Stack with the top item removed
	// along with the item. If the Stack is empty
	// the boolean third return value will be false
	// and the popped value will be nil
	Pop() (Stack, interface{}, bool)

	// Return the top value of the Stack. If the
	// Stack is empty then the return value will be nil
	// and the added boolean flag will be false
	Peek() (interface{}, bool)

	IsEmpty() bool

	Copy() Sequence

	String() string
}

func NewStack() Stack {
	return &EmptyStack{}
}

type EmptyStack struct {
}

func (stack *EmptyStack) IsEmpty() bool {
	return true
}

func (stack *EmptyStack) Size() int {
	return 0
}

func (stack *EmptyStack) Pop() (Stack, interface{}, bool) {
	return stack, nil, false
}

func (stack *EmptyStack) Push(item interface{}) Stack {
	return &NonEmptyStack{
		size: 1,
		head: item,
		tail: stack,
	}
}

func (stack *EmptyStack) String() string {
	return "()"
}

func (stack *EmptyStack) Copy() Sequence {
	return &EmptyStack{}
}

func (stack *EmptyStack) Peek() (interface{}, bool) {
	return nil, false
}

func (stack *EmptyStack) Append(item interface{}) Sequence {
	return stack.Push(item)
}

func (stack *EmptyStack) Get(index int) interface{} {
	panic(ErrIndexOutOfRange)
}

func (stack *EmptyStack) Update(index int, value interface{}) Sequence {
	panic(ErrIndexOutOfRange)
}

func (iterable *EmptyStack) Iterator() Iterator {
	return &EmptyIterator{}
}

func (iterable *EmptyStack) ForEach(iterFn func(interface{})) {

}

func (iterable *EmptyStack) Map(mapFn func(interface{}) interface{}) Iterable {
	return iterable
}

func (iterable *EmptyStack) Filter(filterFn func(interface{}) bool) Iterable {
	return iterable
}

func (iterable *EmptyStack) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	return initialValue
}

func (iterable *EmptyStack) ToSlice() []interface{} {
	return []interface{}{}
}

func (iterable *EmptyStack) Take(count int) Iterable {
	return iterable
}

func (iterable *EmptyStack) Skip(count int) Iterable {
	return iterable
}

func (iterable *EmptyStack) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return iterable
}

func (iterable *EmptyStack) Any(matchFn func(interface{}) bool) bool {
	return false
}

type NonEmptyStack struct {
	size int
	head interface{}
	tail Stack
}

func (stack *NonEmptyStack) IsEmpty() bool {
	return false
}

func (stack *NonEmptyStack) Push(value interface{}) Stack {

	return &NonEmptyStack{
		size: stack.size + 1,
		head: value,
		tail: stack,
	}
}

func (stack *NonEmptyStack) Pop() (Stack, interface{}, bool) {
	return stack.tail, stack.head, true
}

func (stack *NonEmptyStack) Peek() (interface{}, bool) {
	return stack.head, true
}

func (stack *NonEmptyStack) Copy() Sequence {
	return &NonEmptyStack{
		size: stack.size,
		head: stack.head,
		tail: stack.tail.Copy().(Stack),
	}
}

func (stack *NonEmptyStack) Append(item interface{}) Sequence {
	copy := stack.Copy().(*NonEmptyStack)

	lastNonEmptyNode := copy
	lastNonEmptyNode.size += 1
	for !lastNonEmptyNode.tail.IsEmpty() {
		lastNonEmptyNode = lastNonEmptyNode.tail.(*NonEmptyStack)
		lastNonEmptyNode.size += 1
	}

	newTail := NewStack().Push(item)
	lastNonEmptyNode.tail = newTail
	return copy
}

func (stack *NonEmptyStack) Get(index int) interface{} {
	if index == 0 {
		return stack.head
	}
	if index < 0 || index >= stack.size {
		panic(ErrIndexOutOfRange)
	}

	iterator := stack.Iterator()

	currentIndex := 0
	for iterator.MoveNext() {
		if index == currentIndex {
			return iterator.Current()
		}
		currentIndex++
	}
	// Should never happen
	panic(ErrIndexOutOfRange)
}

func (stack *NonEmptyStack) Update(index int, value interface{}) Sequence {
	if index < 0 || index >= stack.size {
		panic(ErrIndexOutOfRange)
	}

	copy := stack.Copy().(*NonEmptyStack)
	if index == 0 {
		copy.head = value
		return copy
	}
	currentIndex := 0
	var current Stack = copy
	for !current.IsEmpty() {
		if index == currentIndex {
			current.(*NonEmptyStack).head = value
		}
		current = current.(*NonEmptyStack).tail
		currentIndex++
	}
	return copy
}

func (stack *NonEmptyStack) Size() int {
	return stack.size
}

func (stack *NonEmptyStack) Iterator() Iterator {
	return &StackIterator{
		initialized: false,
		current:     stack,
	}
}

func (stack *NonEmptyStack) ForEach(iterFn func(interface{})) {
	forEachHelper(stack, iterFn)
}

func (stack *NonEmptyStack) Map(mapFn func(interface{}) interface{}) Iterable {
	return mapHelper(stack, mapFn)
}

func (stack *NonEmptyStack) Filter(filterFn func(interface{}) bool) Iterable {
	return filterHelper(stack, filterFn)
}

func (stack *NonEmptyStack) Fold(initialValue interface{}, reducerFn func(interface{}, interface{}) interface{}) interface{} {
	return foldHelper(stack, initialValue, reducerFn)
}

func (stack *NonEmptyStack) ToSlice() []interface{} {
	return toSliceHelper(stack)
}

func (stack *NonEmptyStack) Take(count int) Iterable {
	return takeHelper(stack, count)
}

func (stack *NonEmptyStack) Skip(count int) Iterable {
	return skipHelper(stack, count)
}

func (stack *NonEmptyStack) SkipWhile(matchFn func(interface{}) bool) Iterable {
	return skipWhileHelper(stack, matchFn)
}

func (stack *NonEmptyStack) Any(matchFn func(interface{}) bool) bool {
	return anyHelper(stack, matchFn)
}

func (stack *NonEmptyStack) String() string {
	return fmt.Sprintf("%v::%v", stack.head, stack.tail.String())
}

type StackIterator struct {
	initialized bool
	current     *NonEmptyStack
}

func (iterator *StackIterator) MoveNext() bool {
	// Workaround to handle the fact that iteration
	// should always start with a call to MoveNext before
	// anyone calls Current
	if !iterator.initialized {
		iterator.initialized = true
		return true
	} else if iterator.current.tail.IsEmpty() {
		return false
	} else {
		iterator.current = iterator.current.tail.(*NonEmptyStack)
		return true
	}
}

func (iterator *StackIterator) Current() interface{} {
	if !iterator.initialized {
		panic(ErrIterationOutOfRange)
	} else {
		return iterator.current.head
	}
}
