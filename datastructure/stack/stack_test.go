package stack_test

import (
	"testing"

	stackPackage "github.com/JudicaelT/betterstandards/datastructure/stack"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkIsEmpty(b *testing.B) {
	stack := stackPackage.NewStack(9, 10, 21)
	codeUnderTest := func() { stack.IsEmpty() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkToSlice(b *testing.B) {
	stack := stackPackage.NewStack(9, 10, 21)
	codeUnderTest := func() { stack.ToSlice() }
	benchmark.AvgRuntime(b, codeUnderTest)
	// There should be 1 alloc because the slice is created
	// with a dynamic length using make([]T, stack.Len())
	benchmark.AssertAvgAllocs(b, 1, codeUnderTest)
}

func BenchmarkLen(b *testing.B) {
	stack := stackPackage.NewStack(9, 10, 21)
	codeUnderTest := func() { stack.Len() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkPush(b *testing.B) {
	stack := stackPackage.NewStack(9, 10)
	codeUnderTest := func() { stack.Push(21) }
	benchmark.AvgRuntime(b, codeUnderTest)
	// There should be 1 alloc because pushing an item
	// appends a new node to the internal list
	benchmark.AssertAvgAllocs(b, 1, codeUnderTest)
}

func BenchmarkPop(b *testing.B) {
	stack := stackPackage.NewStack(9, 10, 21)
	codeUnderTest := func() { stack.Pop() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkMustPop(b *testing.B) {
	stack := stackPackage.NewStack(9, 10, 21)
	codeUnderTest := func() { stack.MustPop() }
	resetStack := func() { stack.Push(21) }
	benchmark.AvgRuntime(b, codeUnderTest, resetStack)
	benchmark.AssertNoAllocs(b, codeUnderTest, resetStack)
}

func TestIsEmpty(t *testing.T) {
	// Given a Stack with items in it
	stack := stackPackage.NewStack(9, 10, 21)
	// When we check if the stack is empty
	var isEmpty bool = stack.IsEmpty()
	// Then it should return false
	assert.False(t, isEmpty)
}

func TestIsEmptyWithEmptyStack(t *testing.T) {
	// Given an empty Stack
	stack := stackPackage.NewStack[any]()
	// When we check if the stack is empty
	var isEmpty bool = stack.IsEmpty()
	// Then it should return true
	assert.True(t, isEmpty)
}

func TestPush(t *testing.T) {
	// Given a Stack
	stack := stackPackage.NewStack(9, 10)
	var length int = stack.Len()
	assert.Equal(t, 2, length)

	// When we push an item
	stack.Push(21)

	// Then the stack's length should have increased by 1
	length = stack.Len()
	assert.Equal(t, 3, length)

	// And the last item should be the pushed item
	var stackSlice []int = stack.ToSlice()
	var pushedItem int = stackSlice[2]
	assert.Equal(t, 21, pushedItem)
}

func TestPop(t *testing.T) {
	// Given a Stack
	stack := stackPackage.NewStack(9, 10, 21)
	var length int = stack.Len()
	assert.Equal(t, 3, length)

	// When we pop
	var poppedItem int
	poppedItem, err := stack.Pop()

	// Then the stack's length should have decreased by 1
	var stackSlice []int = stack.ToSlice()
	length = stack.Len()
	assert.Equal(t, 2, length)

	// And the last inserted item should be removed
	assert.Equal(t, 9, stackSlice[0])
	assert.Equal(t, 10, stackSlice[1])

	// And the popped item should be the last inserted item
	assert.Equal(t, 21, poppedItem)

	// And there should be no error
	assert.NoError(t, err)
}

func TestPopWithEmptyStack(t *testing.T) {
	// Given a Stack
	stack := stackPackage.NewStack[int8]()
	// When we pop
	_, err := stack.Pop()
	// Then there should be an error
	assert.Same(t, stackPackage.PopFromEmptyStackErr, err)
}

func TestMustPop(t *testing.T) {
	// Given a Stack
	stack := stackPackage.NewStack(9, 10, 21)
	var length int = stack.Len()
	assert.Equal(t, 3, length)

	// When we pop
	var poppedItem int = stack.MustPop()

	// Then the stack's length should have decreased by 1
	var stackSlice []int = stack.ToSlice()
	length = stack.Len()
	assert.Equal(t, 2, length)

	// And the last inserted item should be removed
	assert.Equal(t, 9, stackSlice[0])
	assert.Equal(t, 10, stackSlice[1])

	// And the popped items should be the last inserted item
	assert.Equal(t, 21, poppedItem)
}

func TestMustPopWithEmptyStack(t *testing.T) {
	// Given an empty Stack
	stack := stackPackage.NewStack[int]()

	// Stack.MustPop should panic
	functionUnderTest := "Stack.MustPop"
	failMessage := "Tried to pop an item from an empty stack"
	defer test.ShouldPanic(t, functionUnderTest, failMessage)

	// When we try to pop an item from an empty stack
	stack.MustPop()
}

func TestToSlice(t *testing.T) {
	// Given a Stack
	stack := stackPackage.NewStack(9, 10, 21)

	// When we convert the stack into a slice
	var slice []int = stack.ToSlice()

	// Then the created slice should match the stack
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 9, slice[0])
	assert.Equal(t, 10, slice[1])
	assert.Equal(t, 21, slice[2])
}
