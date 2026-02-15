package assert_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertCondition(b *testing.B) {
	codeUnderTest := func() { assert.Condition(true, "fail message") }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertConditionWithDynamicMessage(b *testing.B) {
	codeUnderTest := func() { assert.Condition(true, "%s fail message", "dynamic") }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertCondition(t *testing.T) {
	// assert.Condition() should not panic
	functionUnderTest := "assert.Condition"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When a true condition is passed
	assert.Condition(
		2+2 == 4,
		"Some message if the condition was false",
	)
}

func TestAssertConditionWithFalse(t *testing.T) {

	// assert.Condition() should panic
	functionUnderTest := "assert.Condition"
	failMessage := "Failed asserting that 9 + 10 equals 21"
	defer test.ShouldPanic(t, functionUnderTest, failMessage)

	// When a false condition is passed
	assert.Condition(9+10 == 21, failMessage)
}

func TestAssertConditionWithDynamicMessage(t *testing.T) {
	// assert.Condition() should panic
	failMessage := "This is a %s message with %d parameters"
	param1 := "dynamic"
	param2 := 2

	functionUnderTest := "assert.Condition"
	expectedMessage := fmt.Sprintf(failMessage, param1, param2)
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When a false condition is passed with a dynamic message
	assert.Condition(
		9+10 == 21,
		failMessage,
		param1,
		param2,
	)
}
