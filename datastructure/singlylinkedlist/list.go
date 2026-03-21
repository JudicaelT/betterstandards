package singlylinkedlist

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
)

var InsertAfterOrphanNodeErr error = errors.New("Cannot insert a value after orphan node")

type List[T any] struct {
	head *Node[T]
	len  int
}

func New[T any](values ...T) (list *List[T], tail *Node[T]) {
	list = &List[T]{}
	if 0 != len(values) {
		tail = list.Append(values[0])
		if 2 <= len(values) {
			for _, value := range values[1:] {
				tail = assert.Must(tail.InsertNext(value))
			}
		}
	}
	return list, tail
}

func (l *List[T]) PopHead() *Node[T] {
	var head *Node[T] = l.Head()
	if head != nil {
		head.Remove()
	}
	return head
}

func (l *List[T]) PopTail() *Node[T] {
	var tail, prev *Node[T]
	for tail = l.Head(); tail != nil && tail.Next() != nil; tail = tail.Next() {
		prev = tail
	}
	if prev != nil {
		prev.RemoveNext()
	} else if tail != nil {
		tail.Remove()
	}
	return tail
}

func (l *List[T]) Append(value T) *Node[T] {
	if l.Head() == nil {
		return l.setHead(value)
	}
	return assert.Must(l.Tail().InsertNext(value))
}

func (l *List[T]) Prepend(value T) *Node[T] {
	return l.setHead(value)
}

func (l *List[T]) setHead(value T) *Node[T] {
	newHead := &Node[T]{
		Value: value,
		list:  l,
	}
	if l.Head() != nil {
		newHead.next = l.Head()
	}
	l.head = newHead
	l.len++
	return newHead
}

func (l *List[T]) Tail() *Node[T] {
	var tail *Node[T]
	for tail = l.Head(); tail != nil && tail.Next() != nil; tail = tail.Next() {
	}
	return tail
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Len() int {
	return l.len
}

type Node[T any] struct {
	Value T
	next  *Node[T]
	list  *List[T]
}

func (n *Node[T]) Remove() *Node[T] {
	if n.IsOrphan() {
		return n
	}

	prev := n.findPreviousNode()
	if prev != nil {
		prev.next = n.next
	}

	if n.IsHeadOfList() {
		n.list.head = n.next
	}
	n.list.len--
	n.next = nil
	n.list = nil
	return n
}

func (n *Node[T]) RemoveNext() *Node[T] {
	next := n.Next()
	if next != nil {
		n.list.len--
		n.next = next.Next()
		next.next = nil
		next.list = nil
	}
	return next
}

func (n *Node[T]) findPreviousNode() *Node[T] {
	if n.IsHeadOfList() || n.IsOrphan() {
		return nil
	}

	var previousNode *Node[T]
	for node := n.list.Head(); node != nil && node != n; node = node.Next() {
		previousNode = node
	}
	return previousNode
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) InsertNext(value T) (*Node[T], error) {
	if n.IsOrphan() {
		return nil, InsertAfterOrphanNodeErr
	}
	node := &Node[T]{
		Value: value,
		list:  n.list,
		next:  n.next,
	}
	n.next = node
	n.list.len++
	return node, nil
}

func (n *Node[T]) IsHeadOfList() bool {
	return !n.IsOrphan() && n == n.list.Head()
}

func (n *Node[T]) IsTailOfList() bool {
	return !n.IsOrphan() && n.next == nil
}

func (n *Node[T]) IsOrphan() bool {
	return n.list == nil
}
