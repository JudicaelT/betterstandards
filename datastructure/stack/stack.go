package stack

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/datastructure/doublylinkedlist"
)

var PopFromEmptyStackErr error = errors.New("Tried to pop an item from an empty stack")

type Stack[T any] struct {
	list *doublylinkedlist.List[T]
}

func NewStack[T any](items ...T) *Stack[T] {
	return &Stack[T]{doublylinkedlist.New(items...)}
}

func (q *Stack[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Stack[T]) ToSlice() []T {
	slice := make([]T, q.Len())
	var i int
	for node := q.list.Head(); node != nil; node = node.Next() {
		slice[i] = node.Value
		i++
	}
	return slice
}

func (q *Stack[T]) Len() int {
	return q.list.Len()
}

func (q *Stack[T]) Push(item T) {
	q.list.Append(item)
}

func (q *Stack[T]) MustPop() T {
	return assert.Must(q.Pop())
}

func (q *Stack[T]) Pop() (T, error) {
	tail := q.list.PopTail()
	if tail == nil {
		var item T
		return item, PopFromEmptyStackErr
	}
	return tail.Value, nil
}
