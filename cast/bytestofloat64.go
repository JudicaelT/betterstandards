package cast

import (
	"encoding/binary"
	"fmt"
	"math"
)

func BytesToFloat64(value []byte, byteOrder binary.ByteOrder) (float64, error) {
	if len(value) < 8 {
		return 0, fmt.Errorf(
			"Cannot convert []byte to float64 because it contains less than 8 bytes (%d given)",
			len(value),
		)
	}
	return math.Float64frombits(byteOrder.Uint64(value)), nil
}
