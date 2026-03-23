package cast

import (
	"encoding/binary"
	"fmt"
	"math"
)

type CastBytesToFloat64Err struct {
	BytesProvided int
}

func (c CastBytesToFloat64Err) Error() string {
	return fmt.Sprintf(
		"Cannot convert []byte to float64 because it contains less than 8 bytes (%d given)",
		c.BytesProvided,
	)
}

func BytesToFloat64(bytes []byte, byteOrder binary.ByteOrder) (float64, error) {
	var bytesProvided int = len(bytes)
	if bytesProvided < 8 {
		return 0, CastBytesToFloat64Err{BytesProvided: bytesProvided}
	}
	return math.Float64frombits(byteOrder.Uint64(bytes)), nil
}
