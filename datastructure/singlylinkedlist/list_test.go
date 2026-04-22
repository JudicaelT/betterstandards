package singlylinkedlist_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/datastructure/singlylinkedlist"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkIsEmpty(bench *testing.B) {
	list, _ := singlylinkedlist.New("Hello", "world")
	codeUnderTest := func() { list.IsEmpty() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkLen(bench *testing.B) {
	list, _ := singlylinkedlist.New("Hello", "world")
	codeUnderTest := func() { list.Len() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkAppend(bench *testing.B) {
	list, _ := singlylinkedlist.New("Hello", "world")
	codeUnderTest := func() { list.Append("!") }
	resetList := func() { list.PopTail() }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertAvgAllocs(bench, 1, codeUnderTest, resetList)
}

func BenchmarkAppendToEmptyList(bench *testing.B) {
	list, _ := singlylinkedlist.New[string]()
	codeUnderTest := func() { list.Append("!") }
	resetList := func() { list.PopTail() }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertAvgAllocs(bench, 1, codeUnderTest, resetList)
}

func BenchmarkPrepend(bench *testing.B) {
	list, _ := singlylinkedlist.New("world", "!")
	codeUnderTest := func() { list.Prepend("Hello") }
	resetList := func() { list.PopHead() }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertAvgAllocs(bench, 1, codeUnderTest, resetList)
}

func BenchmarkPrependToEmptyList(bench *testing.B) {
	list, _ := singlylinkedlist.New[string]()
	codeUnderTest := func() { list.Prepend("Hello") }
	resetList := func() { list.PopHead() }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertAvgAllocs(bench, 1, codeUnderTest, resetList)
}

func BenchmarkTraverse(bench *testing.B) {
	list, _ := singlylinkedlist.New(9, 10, 21)
	codeUnderTest := func() {
		for node := list.Head(); node != nil; node = node.Next() {
		}
	}
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkPopHead(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.PopHead() }
	resetList := func() { list.Prepend(9) }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertNoAllocs(bench, codeUnderTest, resetList)
}

func BenchmarkPopHeadOfEmptyList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8]()
	codeUnderTest := func() { list.PopHead() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkPopTail(bench *testing.B) {
	list, _ := singlylinkedlist.New(9, 10, 21)
	codeUnderTest := func() { list.PopTail() }
	resetList := func() { list.Append(21) }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertNoAllocs(bench, codeUnderTest, resetList)
}

func BenchmarkPopTailOfEmptyList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8]()
	codeUnderTest := func() { list.PopTail() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestIsEmpty(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New("Hello", "world")
	// When we check if the list is empty
	var isEmpty bool = list.IsEmpty()
	// Then it should return false
	assert.False(t, isEmpty)
}

func TestEmptyListIsEmpty(t *testing.T) {
	// Given an empty SinglyLinkedList
	list, _ := singlylinkedlist.New[int8]()
	// When we check if the list is empty
	var isEmpty bool = list.IsEmpty()
	// Then it should return true
	assert.True(t, isEmpty)
}

func TestLen(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New("Hello", "world")
	// When we get the length of the list
	var listLen int = list.Len()
	// Then it's length should match the number of elements in the list
	assert.Equal(t, 2, listLen)
}

func TestAppend(t *testing.T) {
	// Given a SinglyLinkedList
	list, oldTail := singlylinkedlist.New("Hello", "world")

	// When we append an item to the list
	var insertedNode *singlylinkedlist.Node[string] = list.Append("!")

	// Then a new node containing the value should be inserted at the end
	expectedList, _ := singlylinkedlist.New("Hello", "world", "!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, oldTail.Next())
}

func TestAppendToEmptyList(t *testing.T) {
	// Given an empty SinglyLinkedList
	list, _ := singlylinkedlist.New[string]()

	// When we append an item to the list
	var insertedNode *singlylinkedlist.Node[string] = list.Append("!")

	// Then a new node containing the value should be inserted and be the head and tail of list
	expectedList, _ := singlylinkedlist.New("!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, list.Head())
	assert.Same(t, insertedNode, list.Tail())
}

func TestPrepend(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New("world", "!")

	// When we prepend an item to the list
	var newHead *singlylinkedlist.Node[string] = list.Prepend("Hello")

	// Then a new node containing the value should be inserted at the start
	expectedList, _ := singlylinkedlist.New("Hello", "world", "!")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, newHead, list.Head())
}

func TestPrependToEmptyList(t *testing.T) {
	// Given an empty SinglyLinkedList
	list, _ := singlylinkedlist.New[string]()

	// When we prepend an item to the list
	var insertedNode *singlylinkedlist.Node[string] = list.Prepend("Hello")

	// Then the inserted item should be the new head and tail of the list
	expectedList, _ := singlylinkedlist.New("Hello")
	assertListsMatch(t, expectedList, list)
	assert.Same(t, insertedNode, list.Head())
	assert.Same(t, insertedNode, list.Tail())
}

func TestTraverse(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New(9, 10, 21)

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

func TestPopHead(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When we pop the head of the list
	var removedHead *singlylinkedlist.Node[int8] = list.PopHead()

	// Then it should be removed from the list
	assert.Equal(t, int8(9), removedHead.Value)
	assertNodeHasBeenRemoved(t, removedHead)
	expectedList, _ := singlylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
	// And the new head should be the node containing 10
	assert.Equal(t, int8(10), list.Head().Value)
}

func TestPopHeadOfEmptyList(t *testing.T) {
	// Given an empty SinglyLinkedList
	list, _ := singlylinkedlist.New[int8]()
	// When we pop the head of the list
	var removedHead *singlylinkedlist.Node[int8] = list.PopHead()
	// Then it should return a nil node
	assert.Nil(t, removedHead)
}

func TestPopTail(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When we pop the tail of the list
	var removedTail *singlylinkedlist.Node[int8] = list.PopTail()

	// Then it should be removed from the list
	assert.Equal(t, int8(21), removedTail.Value)
	assertNodeHasBeenRemoved(t, removedTail)
	expectedList, _ := singlylinkedlist.New[int8](9, 10)
	assertListsMatch(t, expectedList, list)
}

func TestPopTailOfEmptyList(t *testing.T) {
	// Given an empty SinglyLinkedList
	list, _ := singlylinkedlist.New[int8]()
	// When we pop the tail of the list
	var removedTail *singlylinkedlist.Node[int8] = list.PopTail()
	// Then it should return a nil node
	assert.Nil(t, removedTail)
}

func TestPopTailWithOnlyOneElement(t *testing.T) {
	// Given a SinglyLinkedList containing only one element
	list, _ := singlylinkedlist.New[int8](42)

	// When we pop the tail of the list
	var removedTail *singlylinkedlist.Node[int8] = list.PopTail()

	// Then the list should be empty
	assert.Equal(t, int8(42), removedTail.Value)
	assertNodeHasBeenRemoved(t, removedTail)
	expectedList, _ := singlylinkedlist.New[int8]()
	assertListsMatch(t, expectedList, list)
	// And should no longer have a head or tail
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}
