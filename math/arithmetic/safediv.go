package arithmetic

import (
	"errors"
	"math"
	"reflect"

	"github.com/JudicaelT/betterstandards/types"
)

var DivOverflowErr error = errors.New("An overflow occurred while dividing two numbers")
var DivByZeroErr error = errors.New("Cannot divide by zero")

func SafeDiv[T types.Numeric](a, b T, moreNumbersToDiv ...T) (T, error) {
	if b == 0 {
		return 0, DivByZeroErr
	}
	var quotient T = a / b
	var hasOverflowed bool = divHasOverflowed(a, b)
	for _, number := range moreNumbersToDiv {
		if number == 0 {
			return 0, DivByZeroErr
		}
		quotient = quotient / number
		hasOverflowed = hasOverflowed || divHasOverflowed(quotient, number)
	}
	if hasOverflowed {
		return quotient, DivOverflowErr
	}
	return quotient, nil
}

func divHasOverflowed[T types.Numeric](a, b T) bool {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		aInt, _ := any(a).(int)
		bInt, _ := any(b).(int)
		return aInt == math.MinInt && bInt == -1
	case reflect.Int8:
		aInt8, _ := any(a).(int8)
		bInt8, _ := any(b).(int8)
		return aInt8 == math.MinInt8 && bInt8 == -1
	case reflect.Int16:
		aInt16, _ := any(a).(int16)
		bInt16, _ := any(b).(int16)
		return aInt16 == math.MinInt16 && bInt16 == -1
	case reflect.Int32:
		aInt32, _ := any(a).(int32)
		bInt32, _ := any(b).(int32)
		return aInt32 == math.MinInt32 && bInt32 == -1
	case reflect.Int64:
		aInt64, _ := any(a).(int64)
		bInt64, _ := any(b).(int64)
		return aInt64 == math.MinInt64 && bInt64 == -1
	default:
		return false
	}
}
