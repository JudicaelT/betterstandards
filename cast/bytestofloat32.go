package cast

import (
	"encoding/binary"
	"fmt"
	"math"
)

func BytesToFloat32(value []byte, byteOrder binary.ByteOrder) (float32, error) {
	if len(value) < 4 {
		return 0, fmt.Errorf(
			"Cannot convert []byte to float32 because it contains less than 4 bytes (%d given)",
			len(value),
		)
	}
	return math.Float32frombits(byteOrder.Uint32(value)), nil
}
