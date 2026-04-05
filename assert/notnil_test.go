package assert_test

import (
	"testing"
	"unsafe"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	testify "github.com/stretchr/testify/assert"
)

func BenchmarkAssertNotNil(b *testing.B) {
	codeUnderTest := func() { assert.NotNil("Hello world!") }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNotNil(t *testing.T) {
	testify.NotPanics(t, func() { assert.NotNil(struct{}{}) })
	testify.NotPanics(t, func() { assert.NotNil(false) })
	testify.NotPanics(t, func() { assert.NotNil(byte(0)) })
	testify.NotPanics(t, func() { assert.NotNil(int(0)) })
	testify.NotPanics(t, func() { assert.NotNil(int8(0)) })
	testify.NotPanics(t, func() { assert.NotNil(int16(0)) })
	testify.NotPanics(t, func() { assert.NotNil(int32(0)) })
	testify.NotPanics(t, func() { assert.NotNil(int64(0)) })
	testify.NotPanics(t, func() { assert.NotNil(uint(0)) })
	testify.NotPanics(t, func() { assert.NotNil(uint8(0)) })
	testify.NotPanics(t, func() { assert.NotNil(uint16(0)) })
	testify.NotPanics(t, func() { assert.NotNil(uint32(0)) })
	testify.NotPanics(t, func() { assert.NotNil(uint64(0)) })
	testify.NotPanics(t, func() { assert.NotNil(float32(0)) })
	testify.NotPanics(t, func() { assert.NotNil(float64(0)) })
	testify.NotPanics(t, func() { assert.NotNil(complex64(0)) })
	testify.NotPanics(t, func() { assert.NotNil(complex128(0)) })
	testify.NotPanics(t, func() { assert.NotNil("test string") })
	testify.NotPanics(t, func() { assert.NotNil(make(chan any)) })
	testify.NotPanics(t, func() { assert.NotNil(func() {}) })
	testify.NotPanics(t, func() { assert.NotNil(any(42)) })
	testify.NotPanics(t, func() { assert.NotNil(make(map[any]any)) })
	testify.NotPanics(t, func() { assert.NotNil(&struct{}{}) })
	testify.NotPanics(t, func() { assert.NotNil([]any{}) })
	testify.NotPanics(t, func() { assert.NotNil(unsafe.Pointer(&struct{}{})) })
}

func TestAssertNotNilWithNil(t *testing.T) {
	testify.Panics(t, func() { assert.NotNil(nil) })
	testify.Panics(t, func() {
		var uninitializedChan chan any
		assert.NotNil(uninitializedChan)
	})
	testify.Panics(t, func() {
		var uninitializedFunc func()
		assert.NotNil(uninitializedFunc)
	})
	testify.Panics(t, func() {
		var uninitializedInterface any
		assert.NotNil(uninitializedInterface)
	})
	testify.Panics(t, func() {
		var uninitializedMap map[any]any
		assert.NotNil(uninitializedMap)
	})
	testify.Panics(t, func() {
		var uninitializedPointer *string
		assert.NotNil(uninitializedPointer)
	})
	testify.Panics(t, func() {
		var uninitializedSlice []any
		assert.NotNil(uninitializedSlice)
	})
	testify.Panics(t, func() {
		var uninitializedUnsafePointer unsafe.Pointer
		assert.NotNil(uninitializedUnsafePointer)
	})
}
