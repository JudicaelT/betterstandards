package queue_test

import (
	"testing"

	queuePackage "github.com/JudicaelT/betterstandards/datastructure/queue"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkIsEmpty(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10, 21)
	codeUnderTest := func() { queue.IsEmpty() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkToSlice(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10, 21)
	codeUnderTest := func() { queue.ToSlice() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertAvgAllocs(b, 0, codeUnderTest)
}

func BenchmarkLen(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10, 21)
	codeUnderTest := func() { queue.Len() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkEnqueue(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10)
	codeUnderTest := func() { queue.Enqueue(21) }
	benchmark.AvgRuntime(b, codeUnderTest)
	// There should be 1 alloc because enqueuing an item
	// appends a new node to the internal list
	benchmark.AssertAvgAllocs(b, 1, codeUnderTest)
}

func BenchmarkDequeue(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10, 21)
	codeUnderTest := func() { queue.Dequeue() }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkMustDequeue(b *testing.B) {
	queue := queuePackage.NewQueue(9, 10, 21)
	codeUnderTest := func() { queue.MustDequeue() }
	resetQueue := func() { queue.Enqueue(42) }
	benchmark.AvgRuntime(b, codeUnderTest, resetQueue)
	benchmark.AssertNoAllocs(b, codeUnderTest, resetQueue)
}

func TestIsEmpty(t *testing.T) {
	// Given a Queue with items in it
	queue := queuePackage.NewQueue(9, 10, 21)
	// When we check if the queue is empty
	var isEmpty bool = queue.IsEmpty()
	// Then it should return false
	assert.False(t, isEmpty)
}

func TestIsEmptyWithEmptyQueue(t *testing.T) {
	// Given an empty Queue
	queue := queuePackage.NewQueue[any]()
	// When we check if the queue is empty
	var isEmpty bool = queue.IsEmpty()
	// Then it should return true
	assert.True(t, isEmpty)
}

func TestEnqueue(t *testing.T) {
	// Given a Queue
	queue := queuePackage.NewQueue(9, 10)
	var length int = queue.Len()
	assert.Equal(t, 2, length)

	// When we enqueue an item
	queue.Enqueue(21)

	// Then the queue's length should have increased by 1
	var queueSlice []int = queue.ToSlice()
	length = queue.Len()
	assert.Equal(t, 3, length)
	// And the last item should be the enqueued item
	var enqueuedItem int = queueSlice[2]
	assert.Equal(t, 21, enqueuedItem)
}

func TestEnqueueWithEmptyList(t *testing.T) {
	// Given an empty Queue
	queue := queuePackage.NewQueue[int]()
	var length int = queue.Len()
	assert.Equal(t, 0, length)

	// When we enqueue an item
	queue.Enqueue(21)

	// Then the queue's length should have increased by 1
	length = queue.Len()
	assert.Equal(t, 1, length)
}

func TestDequeue(t *testing.T) {
	// Given a Queue
	queue := queuePackage.NewQueue(9, 10, 21)
	var length int = queue.Len()
	assert.Equal(t, 3, length)

	// When we dequeue
	var dequeuedItem int
	dequeuedItem, err := queue.Dequeue()

	// Then the queue's length should have decreased by 1
	var queueSlice []int = queue.ToSlice()
	length = queue.Len()
	assert.Equal(t, 2, length)
	// And the first inserted item should be removed
	assert.Equal(t, 10, queueSlice[0])
	assert.Equal(t, 21, queueSlice[1])
	// And the dequeued item should be the first inserted item
	assert.Equal(t, 9, dequeuedItem)
	// And there should be no error
	assert.NoError(t, err)
}

func TestDequeueWithEmptyQueue(t *testing.T) {
	// Given a Queue
	queue := queuePackage.NewQueue[int8]()
	// When we dequeue
	_, err := queue.Dequeue()
	// Then there should be an error
	assert.Same(t, queuePackage.DequeueFromEmptyQueueErr, err)
}

func TestMustDequeue(t *testing.T) {
	// Given a Queue
	queue := queuePackage.NewQueue(9)
	var length int = queue.Len()
	assert.Equal(t, 1, length)

	// When we dequeue
	var dequeuedItem int = queue.MustDequeue()

	// Then the queue's length should have decreased by 1
	length = queue.Len()
	assert.Equal(t, 0, length)
	// And the dequeued item should be the first inserted item
	assert.Equal(t, 9, dequeuedItem)
}

func TestMustDequeueWithEmptyQueue(t *testing.T) {
	// Given an empty Queue
	queue := queuePackage.NewQueue[int]()

	// Queue.MustDequeue should panic
	functionUnderTest := "Queue.MustDequeue"
	failMessage := "Tried to dequeue an item from an empty queue"
	defer test.ShouldPanic(t, functionUnderTest, failMessage)

	// When we try to dequeue an item from an empty queue
	queue.MustDequeue()
}

func TestToSlice(t *testing.T) {
	// Given a Queue
	queue := queuePackage.NewQueue(9, 10, 21)

	// When we convert the queue into a slice
	var slice []int = queue.ToSlice()

	// Then the created slice should match the queue
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 9, slice[0])
	assert.Equal(t, 10, slice[1])
	assert.Equal(t, 21, slice[2])
}
