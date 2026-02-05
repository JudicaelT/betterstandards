package cast

import (
	"encoding/binary"
	"fmt"
)

func BytesToInt16(value []byte, byteOrder binary.ByteOrder) (int16, error) {
	if len(value) < 2 {
		return 0, fmt.Errorf(
			"Cannot convert []byte to int16 because it contains less than 2 bytes (%d given)",
			len(value),
		)
	}
	return int16(byteOrder.Uint16(value)), nil
}
