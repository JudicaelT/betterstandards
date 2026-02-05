package cast

import (
	"encoding/binary"
	"fmt"
)

func BytesToInt32(value []byte, byteOrder binary.ByteOrder) (int32, error) {
	if len(value) < 4 {
		return 0, fmt.Errorf(
			"Cannot convert []byte to int32 because it contains less than 4 bytes (%d given)",
			len(value),
		)
	}
	return int32(byteOrder.Uint32(value)), nil
}
