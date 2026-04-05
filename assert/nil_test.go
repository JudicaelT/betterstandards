package assert_test

import (
	"errors"
	"testing"
	"unsafe"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	testify "github.com/stretchr/testify/assert"
)

func BenchmarkAssertNil(b *testing.B) {
	codeUnderTest := func() { assert.Nil(nil) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNil(t *testing.T) {
	testify.NotPanics(t, func() { assert.Nil(nil) })
	testify.NotPanics(t, func() {
		var uninitializedChan chan any
		assert.Nil(uninitializedChan)
	})
	testify.NotPanics(t, func() {
		var uninitializedFunc func()
		assert.Nil(uninitializedFunc)
	})
	testify.NotPanics(t, func() {
		var uninitializedInterface any
		assert.Nil(uninitializedInterface)
	})
	testify.NotPanics(t, func() {
		var uninitializedMap map[any]any
		assert.Nil(uninitializedMap)
	})
	testify.NotPanics(t, func() {
		var uninitializedPointer *string
		assert.Nil(uninitializedPointer)
	})
	testify.NotPanics(t, func() {
		var uninitializedSlice []any
		assert.Nil(uninitializedSlice)
	})
	testify.NotPanics(t, func() {
		var uninitializedUnsafePointer unsafe.Pointer
		assert.Nil(uninitializedUnsafePointer)
	})
}

func TestAssertNilWithNonNil(t *testing.T) {
	testify.Panics(t, func() { assert.Nil(struct{}{}) })
	testify.Panics(t, func() { assert.Nil(false) })
	testify.Panics(t, func() { assert.Nil(byte(0)) })
	testify.Panics(t, func() { assert.Nil(int(0)) })
	testify.Panics(t, func() { assert.Nil(int8(0)) })
	testify.Panics(t, func() { assert.Nil(int16(0)) })
	testify.Panics(t, func() { assert.Nil(int32(0)) })
	testify.Panics(t, func() { assert.Nil(int64(0)) })
	testify.Panics(t, func() { assert.Nil(uint(0)) })
	testify.Panics(t, func() { assert.Nil(uint8(0)) })
	testify.Panics(t, func() { assert.Nil(uint16(0)) })
	testify.Panics(t, func() { assert.Nil(uint32(0)) })
	testify.Panics(t, func() { assert.Nil(uint64(0)) })
	testify.Panics(t, func() { assert.Nil(float32(0)) })
	testify.Panics(t, func() { assert.Nil(float64(0)) })
	testify.Panics(t, func() { assert.Nil(complex64(0)) })
	testify.Panics(t, func() { assert.Nil(complex128(0)) })
	testify.Panics(t, func() { assert.Nil("test string") })
	testify.Panics(t, func() { assert.Nil(make(chan any)) })
	testify.Panics(t, func() { assert.Nil(func() {}) })
	testify.Panics(t, func() { assert.Nil(any(42)) })
	testify.Panics(t, func() { assert.Nil(make(map[any]any)) })
	testify.Panics(t, func() { assert.Nil(&struct{}{}) })
	testify.Panics(t, func() { assert.Nil([]any{}) })
	testify.Panics(t, func() { assert.Nil(unsafe.Pointer(&struct{}{})) })
}

func TestAssertNilWithError(t *testing.T) {
	// Given a value that is an error
	value := errors.New("Goodbye world")

	// assert.Nil() should panic
	functionUnderTest := "assert.Nil"
	expectedMessage := "Expected value to be nil. Got error: '" + value.Error() + "'"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When an error is passed
	assert.Nil(value)
}
