package queue

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/datastructure/singlylinkedlist"
)

var DequeueFromEmptyQueueErr error = errors.New("Tried to dequeue an item from an empty queue")

type Queue[T any] struct {
	list *singlylinkedlist.List[T]
	last *singlylinkedlist.Node[T]
}

func NewQueue[T any](items ...T) *Queue[T] {
	list, tail := singlylinkedlist.New(items...)
	return &Queue[T]{list: list, last: tail}
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
	if q.last != nil {
		q.last = assert.Must(q.last.InsertNext(item))
	} else {
		q.last = q.list.Append(item)
	}
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
	if head == q.last {
		q.last = nil
	}
	return head.Value, nil
}
