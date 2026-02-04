package json_test

import (
	"reflect"
	"testing"

	"github.com/JudicaelT/betterstandards/encoding/json"
)

func TestMustMarshal(t *testing.T) {
	// Given an array of strings
	arr := []string{"element1", "element2", "element3"}

	// json.MustMarshal should not panic
	defer func() {
		r := recover()
		if r != nil {
			t.Error("json.MustMarshal() should not have panicked:", r)
		}
	}()

	// When a valid value is passed
	var marshaledArray []byte = json.MustMarshal(arr)

	// And the array should be marshaled correctly
	expectedResult := "[\"element1\",\"element2\",\"element3\"]"
	actualResult := string(marshaledArray)
	if actualResult != expectedResult {
		t.Errorf(
			"Actual result '%s' does not match with expected '%s'",
			actualResult,
			expectedResult,
		)
	}
}

func TestMustMarshalWithInvalidParameter(t *testing.T) {
	// Given a value that cannot be marshaled
	invalidParam := make(chan int)

	// json.MustMarshal should panic
	defer func() {
		r := recover()
		if r == nil {
			t.Error("json.MustMarshal() should have panicked but did not")
		}

		_, isError := r.(error)
		if !isError {
			t.Errorf(
				"json.MustMarshal() should have returned an error but returned '%s' instead",
				reflect.TypeOf(r),
			)
		}
	}()

	// When the invalid value is passed
	json.MustMarshal(invalidParam)
}
