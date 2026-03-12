package doublylinkedlist_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/datastructure/doublylinkedlist"
	"github.com/stretchr/testify/assert"
)

func TestRemoveFromList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When removing a node in the list
	var removedNode *doublylinkedlist.Node[int8] = list.Head().Next().Remove()
	removedNode.Remove() // making sure we can remove a removed node without issues

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, removedNode)
	expectedList := doublylinkedlist.New[int8](9, 21)
	assertListsMatch(t, expectedList, list)
}

func TestRemoveHeadFromList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When removing the head of the list
	var previousHead *doublylinkedlist.Node[int] = list.Head().Remove()

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, previousHead)
	expectedList := doublylinkedlist.New(10, 21)
	assertListsMatch(t, expectedList, list)
	// And the new head should be the node with value 10
	assert.Equal(t, 10, list.Head().Value)
}

func TestRemoveTailFromList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When removing the tail of the list
	var previousTail *doublylinkedlist.Node[int] = list.Tail().Remove()

	// Then it should no longer be part of the list
	assertNodeHasBeenRemoved(t, previousTail)
	expectedList := doublylinkedlist.New(9, 10)
	assertListsMatch(t, expectedList, list)
	// And the new tail should be the node with value 10
	assert.Equal(t, 10, list.Tail().Value)
}

func TestIsHeadOfList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we check if the head of the list is the head of the list
	var isHeadOfList bool = list.Head().IsHeadOfList()
	// Then it should return true
	assert.True(t, isHeadOfList)
	// Otherwise it should return false
	assert.False(t, list.Tail().IsHeadOfList())
}

func TestRemovedNodeIsHeadOfList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we remove the head then check if it is still the head of the list
	var isHeadOfList bool = list.Head().Remove().IsHeadOfList()
	// Then it should return false
	assert.False(t, isHeadOfList)
}

func TestIsTailOfList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we check if the tail of the list is the tail of the list
	var isTailOfList bool = list.Tail().IsTailOfList()
	// Then it should return true
	assert.True(t, isTailOfList)
	// Otherwise it should return false
	assert.False(t, list.Head().IsTailOfList())
}

func TestRemovedNodeIsTailOfList(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we remove the tail then check if it is still the tail of the list
	var isTailOfList bool = list.Tail().Remove().IsTailOfList()
	// Then it should return false
	assert.False(t, isTailOfList)
}

func TestIsOrphan(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we check if a node belonging to that list is an orphan
	var isOrphan bool = list.Head().IsOrphan()
	// Then it should return false
	assert.False(t, isOrphan)
}

func TestRemovedNodeIsOrphan(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)
	// When we check if a node that does not have a list is an orphan
	var isOrphan bool = list.Head().Remove().IsOrphan()
	// Then it should return true
	assert.True(t, isOrphan)
}

func TestInsertNext(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we insert a value after a node
	var insertedNode *doublylinkedlist.Node[int]
	insertedNode, err := list.Head().InsertNext(42)

	// Then a new node containing the value should be inserted after the fiven node
	assert.Equal(t, 42, insertedNode.Value)
	expectedList := doublylinkedlist.New(9, 42, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should be no error
	assert.NoError(t, err)
}

func TestInsertNextTail(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we insert a value after the tail
	var insertedNode *doublylinkedlist.Node[int]
	insertedNode, err := list.Tail().InsertNext(42)

	// Then a new node containing the value should be inserted after the fiven node
	assert.Equal(t, 42, insertedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 21, 42)
	assertListsMatch(t, expectedList, list)
	// And the inserted node should be the tail of the list
	assert.True(t, insertedNode.IsTailOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestInsertNextOnOrphanNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When we insert a value after a node
	var insertedNode *doublylinkedlist.Node[int8]
	insertedNode, err := list.Head().Remove().InsertNext(42)

	// Then it should return an error
	assert.Same(t, doublylinkedlist.InsertAfterOrphanNodeErr, err)
	// And the list should remain the same (aside from the removed node)
	expectedList := doublylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
	// And the returned value should be nil
	assert.Nil(t, insertedNode)
}

func TestInsertPrev(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we insert a value before a node
	var insertedNode *doublylinkedlist.Node[int]
	insertedNode, err := list.Tail().InsertPrev(42)

	// Then the node's previous node should be the inserted value
	assert.Equal(t, 42, insertedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 42, 21)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestInsertPrevOnHead(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we insert a value before the head
	var insertedNode *doublylinkedlist.Node[int]
	insertedNode, err := list.Head().InsertPrev(42)

	// Then it should no longer be the head
	assert.Equal(t, 42, insertedNode.Value)
	expectedList := doublylinkedlist.New(42, 9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And the inserted node should become the new head
	assert.True(t, insertedNode.IsHeadOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestInsertPrevOnOrphanNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When we insert a value before a node
	var insertedNode *doublylinkedlist.Node[int8]
	insertedNode, err := list.Head().Remove().InsertPrev(42)

	// Then it should return an error
	assert.Same(t, doublylinkedlist.InsertBeforeOrphanNodeErr, err)
	// And the list should remain the same (appart from the removed node)
	expectedList := doublylinkedlist.New[int8](10, 21)
	assertListsMatch(t, expectedList, list)
	// And the returned value should be nil
	assert.Nil(t, insertedNode)
}

func TestMoveBefore(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move a node to be before the tail
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Head().Next().MoveBefore(list.Tail())

	// Then the node should be moved before the tail
	assert.Equal(t, 10, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 21, 10, 42)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveHeadBeforeNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move the head of the list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Head().MoveBefore(list.Tail())

	// Then the node should be moved
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(10, 21, 9, 42)
	assertListsMatch(t, expectedList, list)
	// And the moved node should no longer be the head of the list
	assert.False(t, movedNode.IsHeadOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveTailBeforeNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move the tail of the list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Tail().MoveBefore(list.Head().Next())

	// Then the node should be moved
	assert.Equal(t, 42, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 42, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And the moved node should no longer be the tail of the list
	assert.False(t, movedNode.IsTailOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveBeforeNodeNotInSameList(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)
	otherList := doublylinkedlist.New(2, 40, 42)

	// When we move a node from the list to be before a node from the other list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Head().MoveBefore(otherList.Tail())

	// Then the node should be moved from the list to the other list
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(10, 21, 42)
	assertListsMatch(t, expectedList, list)
	expectedOtherList := doublylinkedlist.New(2, 40, 9, 42)
	assertListsMatch(t, expectedOtherList, otherList)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveBeforeNil(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When pass nil as the parameter
	movedNode, err := list.Head().MoveBefore(nil)

	// Then there should be an error
	assert.Same(t, doublylinkedlist.MoveBeforeNilErr, err)
	// And the list should remain unchanged
	expectedList := doublylinkedlist.New[int8](9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And the returned value should be nil
	assert.Nil(t, movedNode)
}

func TestMoveBeforeSamePosition(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to the same position
	movedNode, err := list.Head().MoveBefore(list.Head().Next())

	// Then the list should remain unchanged
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveBeforeSelf(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to be before itself
	movedNode, err := list.Head().MoveBefore(list.Head())

	// Then the list should remain unchanged
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveBeforeOrphanNode(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to be before an orphan node
	orphanNode := list.Tail().Remove()
	movedNode, err := list.Head().MoveBefore(orphanNode)

	// Then the list should remain unchanged (aside from the removed node)
	assert.Nil(t, movedNode)
	expectedList := doublylinkedlist.New(9, 10)
	assertListsMatch(t, expectedList, list)
	// And there should be an error
	assert.Same(t, doublylinkedlist.MoveBeforeOrphanNodeErr, err)
}

func TestMoveAfter(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move a node to be after the head
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Tail().Prev().MoveAfter(list.Head())

	// Then the node should be moved after the head
	assert.Equal(t, 21, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 21, 10, 42)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveHeadAfterNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move the head of the list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Head().MoveAfter(list.Tail().Prev())

	// Then the new head should be the previous head's next node
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(10, 21, 9, 42)
	assertListsMatch(t, expectedList, list)
	// And the moved node should no longer be the head
	assert.False(t, movedNode.IsHeadOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveTailAfterNode(t *testing.T) {
	// Given a DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)

	// When we move the tail of the list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Tail().MoveAfter(list.Head())

	// Then the new tail should be the previous tail's previous node
	assert.Equal(t, 42, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 42, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And the moved node should no longer be the tail
	assert.False(t, movedNode.IsTailOfList())
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveAfterNodeNotInSameList(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21, 42)
	otherList := doublylinkedlist.New(2, 40, 42)

	// When we move a node to be after a node from another list
	var movedNode *doublylinkedlist.Node[int]
	movedNode, err := list.Head().MoveAfter(otherList.Head().Next())

	// Then the node should be moved from the list to the other list
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(10, 21, 42)
	assertListsMatch(t, expectedList, list)
	expectedOtherList := doublylinkedlist.New(2, 40, 9, 42)
	assertListsMatch(t, expectedOtherList, otherList)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveAfterNil(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New[int8](9, 10, 21)

	// When pass nil as the parameter
	movedNode, err := list.Head().MoveAfter(nil)

	// Then there should be an error
	assert.Same(t, doublylinkedlist.MoveAfterNilErr, err)
	// And the list should remain unchanged
	expectedList := doublylinkedlist.New[int8](9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And the returned value should be nil
	assert.Nil(t, movedNode)
}

func TestMoveAfterSamePosition(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to the same position
	movedNode, err := list.Tail().MoveAfter(list.Head().Next())

	// Then the list should remain unchanged
	assert.Equal(t, 21, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveAfterSelf(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to be after itself
	movedNode, err := list.Head().MoveAfter(list.Head())

	// Then the list should remain unchanged
	assert.Equal(t, 9, movedNode.Value)
	expectedList := doublylinkedlist.New(9, 10, 21)
	assertListsMatch(t, expectedList, list)
	// And there should not be an error
	assert.NoError(t, err)
}

func TestMoveAfterOrphanNode(t *testing.T) {
	// Given two DoublyLinkedList
	list := doublylinkedlist.New(9, 10, 21)

	// When we move the node to be after an orphan node
	orphanNode := list.Tail().Remove()
	movedNode, err := list.Head().MoveAfter(orphanNode)

	// Then the list should remain unchanged (aside from the removed node)
	assert.Nil(t, movedNode)
	expectedList := doublylinkedlist.New(9, 10)
	assertListsMatch(t, expectedList, list)
	// And there should be an error
	assert.Same(t, doublylinkedlist.MoveAfterOrphanNodeErr, err)
}

func assertNodeHasBeenRemoved[T any](t *testing.T, node *doublylinkedlist.Node[T]) {
	assert.True(t, node.IsOrphan())
	assert.Nil(t, node.Prev())
	assert.Nil(t, node.Next())
}

func assertListsMatch[T any](t *testing.T, expected, actual *doublylinkedlist.List[T]) {
	assert.Equal(t, expected.Len(), actual.Len())
	if expected.Head() == nil || expected.Tail() == nil {
		assert.Nil(t, actual.Head())
		assert.Nil(t, actual.Tail())
		return
	}

	if !assert.NotNil(t, actual.Head()) || !assert.NotNil(t, actual.Tail()) {
		return
	}

	assert.Equal(t, expected.Head().Value, actual.Head().Value)
	assert.Equal(t, expected.Tail().Value, actual.Tail().Value)

	assert.Nil(t, actual.Head().Prev())
	assert.Nil(t, actual.Tail().Next())

	if expected.Len() == 1 {
		assert.Nil(t, actual.Head().Next())
		assert.Nil(t, actual.Tail().Prev())
		return
	}

	// Check nodes' next pointers by traversing from head to tail
	expectedNode := expected.Head()
	for actualNode := actual.Head(); actualNode != nil; actualNode = actualNode.Next() {
		assert.Equal(t, expectedNode.Value, actualNode.Value)
		expectedNode = expectedNode.Next()
	}

	// Check nodes' previous pointers by traversing from tail to head
	expectedNode = expected.Tail()
	for actualNode := actual.Tail(); actualNode != nil; actualNode = actualNode.Prev() {
		assert.Equal(t, expectedNode.Value, actualNode.Value)
		expectedNode = expectedNode.Prev()
	}
}
