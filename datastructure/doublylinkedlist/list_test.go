package doublylinkedlist_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/datastructure/doublylinkedlist"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New("Hello", "world")
	// When we check if the list is empty
	var isEmpty bool = list.IsEmpty()
	// Then it should return false
	assert.False(t, isEmpty)
}

func TestEmptyListIsEmpty(t *testing.T) {
	// Given an empty DoublyLinkedList
	list := doublylinkedlist.New[int8]()
	// When we check if the list is empty
	var isEmpty bool = list.IsEmpty()
	// Then it should return true
	assert.True(t, isEmpty)
}

func TestLen(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New("Hello", "world")
	// When we get the length of the list
	var listLen int = list.Len()
	// Then it's length should match the number of elements in the list
	assert.Equal(t, 2, listLen)
}

func TestAppend(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New("Hello", "world")

	// When we append an item to the list
	var insertedNode *doublylinkedlist.Node[string] = list.Append("!")

	// Then a new node containing the value should be inserted at the end
	expectedList := doublylinkedlist.New("Hello", "world", "!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, list.Tail())
}

func TestAppendToEmptyList(t *testing.T) {
	// Given an empty DoublyLinkedList
	list := doublylinkedlist.New[string]()

	// When we append an item to the list
	var insertedNode *doublylinkedlist.Node[string] = list.Append("!")

	// Then a new node containing the value should be inserted and be the head and tail of list
	expectedList := doublylinkedlist.New("!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, list.Head())
	assert.Same(t, insertedNode, list.Tail())
}

func TestPrepend(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New("world", "!")

	// When we prepend an item to the list
	var newHead *doublylinkedlist.Node[string] = list.Prepend("Hello")

	// Then a new node containing the value should be inserted at the start
	expectedList := doublylinkedlist.New("Hello", "world", "!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, newHead, list.Head())
}

func TestPrependToEmptyList(t *testing.T) {
	// Given an empty DoublyLinkedList
	list := doublylinkedlist.New[string]()

	// When we prepend an item to the list
	var insertedNode *doublylinkedlist.Node[string] = list.Prepend("Hello")

	// Then the inserted item should be the new head and tail of the list
	expectedList := doublylinkedlist.New("Hello")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, list.Head())
	assert.Same(t, insertedNode, list.Tail())
}

func TestTraverse(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we traverse the list from head to tail
	var i int
	for node := list.Head(); node != nil; node = node.Next() {
		i++

		// Then the value should match each iteration
		value := node.Value
		switch i {
		case 1:
			assert.Equal(t, 9, value)
		case 2:
			assert.Equal(t, 10, value)
		case 3:
			assert.Equal(t, 21, value)
		default:
			t.Fatalf("i should be equal to 1, 2 or 3. Got: %d", i)
		}
	}

	// And there should have been exactly 3 iterations
	assert.Equal(t, 3, i)
}

func TestTraverseBackward(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we traverse the list from tail to head
	var i int
	for node := list.Tail(); node != nil; node = node.Prev() {
		i++

		// Then the value should match each iteration
		value := node.Value
		switch i {
		case 1:
			assert.Equal(t, 21, value)
		case 2:
			assert.Equal(t, 10, value)
		case 3:
			assert.Equal(t, 9, value)
		default:
			t.Fatalf("i should be equal to 1, 2 or 3. Got: %d", i)
		}
	}

	// And there should have been exactly 3 iterations
	assert.Equal(t, 3, i)
}

func TestPopHead(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When we pop the head of the list
	var removedHead *doublylinkedlist.Node[int8] = list.PopHead()

	// Then it should be removed from the list
	assert.Equal(t, int8(9), removedHead.Value)
	assertNodeHasBeenRemoved(t, removedHead)
	expectedList := doublylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
}

func TestPopHeadOfEmptyList(t *testing.T) {
	// Given an empty DoublyLinkedList
	list := doublylinkedlist.New[int8]()
	// When we pop the head of the list
	var removedHead *doublylinkedlist.Node[int8] = list.PopHead()
	// Then it should return a nil node
	assert.Nil(t, removedHead)
}

func TestPopTail(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we pop the tail of the list
	var removedTail *doublylinkedlist.Node[int] = list.PopTail()

	// Then it should be removed from the list
	assert.Equal(t, 21, removedTail.Value)
	assertNodeHasBeenRemoved(t, removedTail)
	expectedList := doublylinkedlist.New(9, 10)
	assertListsMatch(t, expectedList, list)
}

func TestPopTailOfEmptyList(t *testing.T) {
	// Given an empty DoublyLinkedList
	list := doublylinkedlist.New[int8]()
	// When we pop the tail of the list
	var removedTail *doublylinkedlist.Node[int8] = list.PopTail()
	// Then it should return a nil node
	assert.Nil(t, removedTail)
}
