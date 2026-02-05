package cast

import (
	"encoding/binary"
	"fmt"
)

func BytesToInt64(value []byte, byteOrder binary.ByteOrder) (int64, error) {
	if len(value) < 8 {
		return 0, fmt.Errorf(
			"Cannot convert []byte to int64 because it contains less than 8 bytes (%d given)",
			len(value),
		)
	}
	return int64(byteOrder.Uint64(value)), nil
}
