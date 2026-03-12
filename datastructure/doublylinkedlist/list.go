package doublylinkedlist

import (
	"errors"

	"github.com/JudicaelT/betterstandards/assert"
)

var InsertBeforeOrphanNodeErr error = errors.New("Cannot insert a value before orphan node")
var InsertAfterOrphanNodeErr error = errors.New("Cannot insert a value after orphan node")
var MoveBeforeNilErr error = errors.New("Cannot pass nil to doublylinkedlist.Node.MoveBefore()")
var MoveAfterNilErr error = errors.New("Cannot pass nil to doublylinkedlist.Node.MoveAfter()")
var MoveBeforeOrphanNodeErr error = errors.New("Cannot pass orphan node to doublylinkedlist.Node.MoveBefore()")
var MoveAfterOrphanNodeErr error = errors.New("Cannot pass orphan node to doublylinkedlist.Node.MoveAfter()")

type List[T any] struct {
	root *Node[T]
	len  int
}

func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	root := &Node[T]{list: list}
	root.prev = root
	root.next = root
	list.root = root
	for _, value := range values {
		list.Append(value)
	}
	return list
}

func (l *List[T]) PopHead() *Node[T] {
	var head *Node[T] = l.Head()
	if head != nil {
		head.Remove()
	}
	return head
}

func (l *List[T]) Head() *Node[T] {
	return l.root.Next()
}

func (l *List[T]) PopTail() *Node[T] {
	var tail *Node[T] = l.Tail()
	if tail != nil {
		tail.Remove()
	}
	return tail
}

func (l *List[T]) Tail() *Node[T] {
	return l.root.Prev()
}

func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) Append(value T) *Node[T] {
	return assert.Must(l.root.InsertPrev(value))
}

func (l *List[T]) Prepend(value T) *Node[T] {
	return assert.Must(l.root.InsertNext(value))
}

type Node[T any] struct {
	Value      T
	prev, next *Node[T]
	list       *List[T]
}

func (n *Node[T]) Remove() *Node[T] {
	if n.IsOrphan() {
		return n
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.list.len--
	n.prev = nil
	n.next = nil
	n.list = nil
	return n
}

func (n *Node[T]) Next() *Node[T] {
	if n.list != nil && n.next != n.list.root {
		return n.next
	}
	return nil
}

func (n *Node[T]) Prev() *Node[T] {
	if n.list != nil && n.prev != n.list.root {
		return n.prev
	}
	return nil
}

func (n *Node[T]) InsertNext(value T) (*Node[T], error) {
	if n.IsOrphan() {
		return nil, InsertAfterOrphanNodeErr
	}
	node := &Node[T]{
		Value: value,
		list:  n.list,
		next:  n.next,
		prev:  n,
	}
	n.next.prev = node
	n.next = node
	n.list.len++
	return node, nil
}

func (n *Node[T]) InsertPrev(value T) (*Node[T], error) {
	if n.IsOrphan() {
		return nil, InsertBeforeOrphanNodeErr
	}
	node := &Node[T]{
		Value: value,
		list:  n.list,
		next:  n,
		prev:  n.prev,
	}
	n.prev.next = node
	n.prev = node
	n.list.len++
	return node, nil
}

func (n *Node[T]) MoveBefore(node *Node[T]) (*Node[T], error) {
	if node == nil {
		return nil, MoveBeforeNilErr
	}

	if node.IsOrphan() {
		return nil, MoveBeforeOrphanNodeErr
	}

	if node == n || node == n.Next() {
		return n, nil
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	if n.list != node.list {
		n.list.len--
		n.list = node.list
		node.list.len++
	}

	n.prev = node.prev
	n.next = node
	if node.prev != nil {
		node.prev.next = n
	}
	node.prev = n

	return n, nil
}

func (n *Node[T]) MoveAfter(node *Node[T]) (*Node[T], error) {
	if node == nil {
		return nil, MoveAfterNilErr
	}

	if node.IsOrphan() {
		return nil, MoveAfterOrphanNodeErr
	}

	if node == n || node == n.prev {
		return n, nil
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	if n.list != node.list {
		n.list.len--
		n.list = node.list
		node.list.len++
	}

	n.prev = node
	n.next = node.next

	node.next.prev = n
	node.next = n

	return n, nil
}

func (n *Node[T]) IsHeadOfList() bool {
	return !n.IsOrphan() && n == n.list.Head()
}

func (n *Node[T]) IsTailOfList() bool {
	return !n.IsOrphan() && n == n.list.Tail()
}

func (n *Node[T]) IsOrphan() bool {
	return n.list == nil
}
