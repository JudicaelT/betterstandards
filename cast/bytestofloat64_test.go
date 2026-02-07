package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
)

func TestCastBytesToFloat64LittleEndianWithMoreThanThreeBytes(t *testing.T) {
	// Given an slice of height bytes and an slice of more than height bytes
	heightByteSlice := []byte{0, 0, 0, 0, 0, 0, 69, 64}
	nineByteSlice := []byte{0, 0, 0, 0, 0, 0, 69, 64, 40}

	// When we cast the slices to float64 using LittleEndian
	for _, byteSlice := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToFloat64(byteSlice, binary.LittleEndian)

		// Then it should return
		// math.Float64frombits(64 * 256^7 + 69 * 256^6)
		// = math.Float64frombits(0x4045000000000000) = 42
		if result != 42 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 42. Got %f",
				byteSlice,
				result,
			)
		}

		// And not return an error
		if err != nil {
			t.Errorf(
				"(Dataset: %v) Failed asserting that error nil: %s",
				byteSlice,
				err,
			)
		}
	}
}

func TestCastBytesToFloat64BigEndianWithMoreThanThreeByte(t *testing.T) {
	// Given an slice of height bytes and an slice of more than height bytes
	heightByteSlice := []byte{64, 69, 0, 0, 0, 0, 0, 0}
	nineByteSlice := []byte{64, 69, 0, 0, 0, 0, 0, 0, 64}

	// When we cast the slices to float64 using BigEndian
	for _, byteSlice := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToFloat64(byteSlice, binary.BigEndian)

		// Then it should return
		// math.Float64frombits(64 * 256^7 + 69 * 256^6)
		// = math.Float64frombits(0x4045000000000000) = 42
		if result != 42 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 42 Got %f",
				byteSlice,
				result,
			)
		}

		// And not return an error
		if err != nil {
			t.Errorf(
				"(Dataset: %v) Failed asserting that error nil: %s",
				byteSlice,
				err,
			)
		}
	}
}

func TestCastBytesToFloat64WithLessThanFourBytes(t *testing.T) {
	// Given an slice of seven bytes
	sevenByteSlice := []byte{0, 0, 0, 0, 0, 0, 69}

	// When we cast the slice to an float64
	result, err := cast.BytesToFloat64(sevenByteSlice, binary.LittleEndian)

	// Then it should return 0
	if result != 0 {
		t.Error("Failed asserting that result is equal to 0. Got", result)
	}

	// And also return an error with the expected message
	if err == nil {
		t.Fatalf("Failed asserting that error is not nil")
	}
	expectedMessage := "Cannot convert []byte to float64 because it contains less than 8 bytes (7 given)"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
