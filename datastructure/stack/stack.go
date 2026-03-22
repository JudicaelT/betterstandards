package stack

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/datastructure/singlylinkedlist"
)

var PopFromEmptyStackErr error = errors.New("Tried to pop an item from an empty stack")

type Stack[T any] struct {
	list *singlylinkedlist.List[T]
}

func NewStack[T any](items ...T) *Stack[T] {
	list, _ := singlylinkedlist.New[T]()
	for i := len(items) - 1; i >= 0; i-- {
		list.Append(items[i])
	}
	return &Stack[T]{list: list}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) ToSlice() []T {
	slice := make([]T, s.Len())
	var i int = s.Len()
	for node := s.list.Head(); node != nil; node = node.Next() {
		i--
		slice[i] = node.Value
	}
	return slice
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Push(item T) {
	s.list.Prepend(item)
}

func (s *Stack[T]) MustPop() T {
	return assert.Must(s.Pop())
}

func (s *Stack[T]) Pop() (T, error) {
	head := s.list.PopHead()
	if head == nil {
		var item T
		return item, PopFromEmptyStackErr
	}
	return head.Value, nil
}
