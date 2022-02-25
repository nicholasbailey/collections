package collections

import (
	"testing"
)

func TestStackPushAndPop(t *testing.T) {
	expect := expectFor(t)
	stack0 := NewStack()
	stack1 := stack0.Push(5)
	stack2 := stack1.Push(8)
	stack3, val3, found3 := stack2.Pop()
	stack4, val4, found4 := stack3.Pop()
	stack5, val5, found5 := stack4.Pop()

	expect(stack0.Size()).ToBe(0)
	expect(stack1.Size()).ToBe(1)
	expect(stack2.Size()).ToBe(2)
	expect(stack3.Size()).ToBe(1)
	expect(stack4.Size()).ToBe(0)
	expect(stack5.Size()).ToBe(0)
	expect(val3).ToBe(8)
	expect(found3).ToBe(true)
	expect(val4).ToBe(5)
	expect(found4).ToBe(true)
	expect(val5).ToBe(nil)
	expect(found5).ToBe(false)
}

func TestStackPeek(t *testing.T) {
	expect := expectFor(t)
	stack0 := NewStack()
	val, found := stack0.Peek()
	expect(val).ToBe(nil)
	expect(found).ToBe(false)

	stack1 := stack0.Push(1)
	val1, found1 := stack1.Peek()
	expect(val1).ToBe(1)
	expect(found1).ToBe(true)

	stack2 := stack1.Push(2)
	val2, found2 := stack2.Peek()
	expect(val2).ToBe(2)
	expect(found2).ToBe(true)

	// Make sure we don't mutate
	val, found = stack0.Peek()
	expect(val).ToBe(nil)
	expect(found).ToBe(false)

	val1, found1 = stack1.Peek()
	expect(val1).ToBe(1)
	expect(found1).ToBe(true)
}

func TestStackAppend(t *testing.T) {
	expect := expectFor(t)
	stack1 := NewStack().Push(1).Push(2)
	stack2 := stack1.Append(0).(Stack)

	stack3, val3, found3 := stack2.Pop()
	expect(stack3.Size()).ToBe(2)
	expect(val3).ToBe(2)
	expect(found3).ToBe(true)

	stack4, val4, found4 := stack3.Pop()
	expect(stack4.Size()).ToBe(1)
	expect(val4).ToBe(1)
	expect(found4).ToBe(true)

	val5, found5 := stack4.Peek()

	expect(val5).ToBe(0)
	expect(found5).ToBe(true)
}

func TestStackToSlice(t *testing.T) {
	expect := expectFor(t)
	stack := NewStack().Push(1).Push(2).Push(3).Push(4)

	slice := stack.ToSlice()
	expect(slice).ToDeepEqual([]interface{}{4, 3, 2, 1})
}

func TestStackCopy(t *testing.T) {
	expect := expectFor(t)
	stack := NewStack()
	copy := stack.Copy()
	expect(copy.Size()).ToBe(0)
	// We *would* check for inequality here,
	// but because empty stacks have no size, we
	// can't assume anything about pointers

	stack = stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	copy = stack.Copy()

	expect(stack).Not().ToBe(copy)
	expect(stack.Size()).ToBe(copy.Size())
	expect(stack.ToSlice()).ToDeepEqual(copy.ToSlice())
}

func TestStackGetEmptyStack(t *testing.T) {
	expect := expectFor(t)
	stack := NewStack()
	expect(func() { stack.Get(0) }).ToPanicWith(ErrIndexOutOfRange)
}

func TestStackGet(t *testing.T) {
	expect := expectFor(t)
	stack := NewStack().Push(1).Push(2).Push(3).Push(4)
	expect(func() { stack.Get(-1) }).ToPanicWith(ErrIndexOutOfRange)
	expect(func() { stack.Get(4) }).ToPanicWith(ErrIndexOutOfRange)
	expect(stack.Get(0)).ToBe(4)
	expect(stack.Get(1)).ToBe(3)
	expect(stack.Get(2)).ToBe(2)
	expect(stack.Get(3)).ToBe(1)
}

func TestStackUpdate(t *testing.T) {
	expect := expectFor(t)
	stack := NewStack().Push(1).Push(2).Push(3).Push(4)
	expect(func() { stack.Update(-1, 1) }).ToPanicWith(ErrIndexOutOfRange)
	expect(func() { stack.Update(4, 1) }).ToPanicWith(ErrIndexOutOfRange)

	updated := stack.Update(2, 77)
	expect(updated.ToSlice()).ToDeepEqual([]interface{}{4, 3, 77, 1})
	expect(stack.ToSlice()).ToDeepEqual([]interface{}{4, 3, 2, 1})
}
