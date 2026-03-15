package queue

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/datastructure/doublylinkedlist"
)

var DequeueFromEmptyQueueErr error = errors.New("Tried to dequeue an item from an empty queue")

type Queue[T any] struct {
	list *doublylinkedlist.List[T]
}

func NewQueue[T any](items ...T) *Queue[T] {
	return &Queue[T]{doublylinkedlist.New(items...)}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) ToSlice() []T {
	slice := make([]T, q.Len())
	var i int
	for node := q.list.Head(); node != nil; node = node.Next() {
		slice[i] = node.Value
		i++
	}
	return slice
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Enqueue(item T) {
	q.list.Append(item)
}

func (q *Queue[T]) MustDequeue() T {
	return assert.Must(q.Dequeue())
}

func (q *Queue[T]) Dequeue() (T, error) {
	head := q.list.PopHead()
	if head == nil {
		var item T
		return item, DequeueFromEmptyQueueErr
	}
	return head.Value, nil
}
