package singlylinkedlist_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/datastructure/singlylinkedlist"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkRemoveFromList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Head().Next().Remove() }
	resetList := func() { list.Head().InsertNext(10) }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertNoAllocs(bench, codeUnderTest, resetList)
}

func BenchmarkNext(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Head().Next() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkInsertNext(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Tail().InsertNext(42) }
	resetList := func() { list.PopTail() }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertAvgAllocs(bench, 1, codeUnderTest, resetList)
}

func BenchmarkRemoveHeadFromList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Head().Remove() }
	resetList := func() { list.Prepend(9) }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertNoAllocs(bench, codeUnderTest, resetList)
}

func BenchmarkRemoveNext(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Head().RemoveNext() }
	resetList := func() { list.Head().InsertNext(10) }
	benchmark.AvgRuntime(bench, codeUnderTest, resetList)
	benchmark.AssertNoAllocs(bench, codeUnderTest, resetList)
}

func BenchmarkIsHeadOfList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Head().IsHeadOfList() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkIsTailOfList(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Tail().IsTailOfList() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkIsOrphan(bench *testing.B) {
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	codeUnderTest := func() { list.Tail().IsTailOfList() }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestRemoveFromList(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When removing a node in the list
	var removedNode *singlylinkedlist.Node[int8] = list.Head().Next().Remove()
	removedNode.Remove() // making sure we can remove an orphan node

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, removedNode)
	expectedList, _ := singlylinkedlist.New[int8](9, 21)
	assertListsMatch(t, expectedList, list)
}

func TestRemoveHeadFromList(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When removing the head of the list
	var removedNode *singlylinkedlist.Node[int8] = list.Head().Remove()

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, removedNode)
	expectedList, _ := singlylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
	// And the new head should be the node containing 10
	assert.Equal(t, int8(10), list.Head().Value)
}

func TestRemoveNext(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When removing a node in the list using it's previous node
	var removedNode *singlylinkedlist.Node[int8] = list.Head().RemoveNext()

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, removedNode)
	expectedList, _ := singlylinkedlist.New[int8](9, 21)
	assertListsMatch(t, expectedList, list)
}

func TestRemoveTailNextNode(t *testing.T) {
	// Given a SinglyLinkedList
	list, tail := singlylinkedlist.New[int8](9, 10, 21)

	// When removing the tail's next node
	var removedNode *singlylinkedlist.Node[int8] = tail.RemoveNext()

	// Then it should return nil
	assert.Nil(t, removedNode)
	// And the list should remain unchanged
	expectedList, _ := singlylinkedlist.New[int8](9, 10, 21)
	assertListsMatch(t, expectedList, list)
}

func TestIsHeadOfList(t *testing.T) {
	// Given a SinglyLinkedList
	list, tail := singlylinkedlist.New[int8](9, 10, 21)
	// When we check if the head of the list is the head of the list
	var isHeadOfList bool = list.Head().IsHeadOfList()
	// Then it should return true
	assert.True(t, isHeadOfList)
	// Otherwise it should return false
	assert.False(t, tail.IsHeadOfList())
}

func TestRemovedNodeIsHeadOfList(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	// When we remove the head then check if it is still the head of the list
	var isHeadOfList bool = list.Head().Remove().IsHeadOfList()
	// Then it should return false
	assert.False(t, isHeadOfList)
}

func TestIsTailOfList(t *testing.T) {
	// Given a SinglyLinkedList
	list, tail := singlylinkedlist.New[int8](9, 10, 21)
	// When we check if the tail of the list is the tail of the list
	var isTailOfList bool = tail.IsTailOfList()
	// Then it should return true
	assert.True(t, isTailOfList)
	// Otherwise it should return false
	assert.False(t, list.Head().IsTailOfList())
}

func TestRemovedNodeIsTailOfList(t *testing.T) {
	// Given a SinglyLinkedList
	_, tail := singlylinkedlist.New[int8](9, 10, 21)
	// When we remove the tail then check if it is still the tail of the list
	var isTailOfList bool = tail.Remove().IsTailOfList()
	// Then it should return false
	assert.False(t, isTailOfList)
}

func TestIsOrphan(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	// When we check if a node belonging to that list is an orphan
	var isOrphan bool = list.Head().IsOrphan()
	// Then it should return false
	assert.False(t, isOrphan)
}

func TestRemovedNodeIsOrphan(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)
	// When we check if a node that does not have a list is an orphan
	var isOrphan bool = list.Head().Remove().IsOrphan()
	// Then it should return true
	assert.True(t, isOrphan)
}

func TestInsertNext(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New(9, 10, 21)

	// When we insert a value after a node
	var insertedNode *singlylinkedlist.Node[int]
	insertedNode, err := list.Head().InsertNext(42)

	// Then a new node containing the value should be inserted after the fiven node
	assert.Equal(t, 42, insertedNode.Value)
	expectedList, _ := singlylinkedlist.New(9, 42, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should be no error
	assert.NoError(t, err)
}

func TestInsertNextTail(t *testing.T) {
	// Given a SinglyLinkedList
	list, tail := singlylinkedlist.New(9, 10, 21)

	// When we insert a value after the tail
	var insertedNode *singlylinkedlist.Node[int]
	insertedNode, err := tail.InsertNext(42)

	// Then a new node containing the value should be inserted after the fiven node
	assert.Equal(t, 42, insertedNode.Value)
	expectedList, _ := singlylinkedlist.New(9, 10, 21, 42)
	assertListsMatch(t, expectedList, list)
	// And the inserted node should be the tail of the list
	assert.True(t, insertedNode.IsTailOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestInsertNextOnOrphanNode(t *testing.T) {
	// Given a SinglyLinkedList
	list, _ := singlylinkedlist.New[int8](9, 10, 21)

	// When we insert a value after a node
	var insertedNode *singlylinkedlist.Node[int8]
	insertedNode, err := list.Head().Remove().InsertNext(42)

	// Then it should return an error
	assert.Same(t, singlylinkedlist.InsertAfterOrphanNodeErr, err)
	// And the list should remain the same (aside from the removed node)
	expectedList, _ := singlylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
	// And the returned value should be nil
	assert.Nil(t, insertedNode)
}

func assertNodeHasBeenRemoved[T any](t *testing.T, node *singlylinkedlist.Node[T]) {
	assert.True(t, node.IsOrphan())
	assert.Nil(t, node.Next())
}

func assertListsMatch[T any](t *testing.T, expected, actual *singlylinkedlist.List[T]) {
	assert.Equal(t, expected.Len(), actual.Len())
	if expected.Head() == nil || expected.Len() == 0 {
		assert.Nil(t, actual.Head())
		return
	}

	if !assert.NotNil(t, actual.Head()) {
		return
	}

	assert.Equal(t, expected.Head().Value, actual.Head().Value)

	if expected.Len() == 1 {
		assert.True(t, actual.Head().IsHeadOfList())
		assert.True(t, actual.Head().IsTailOfList())
		return
	}

	// Check pointers by traversing from head to tail
	expectedNode := expected.Head()
	for actualNode := actual.Head(); actualNode != nil; actualNode = actualNode.Next() {
		assert.Equal(t, expectedNode.Value, actualNode.Value)
		expectedNode = expectedNode.Next()
	}
}
