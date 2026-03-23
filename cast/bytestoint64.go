package cast

import (
	"encoding/binary"
	"fmt"
)

type CastBytesToInt64Err struct {
	BytesProvided int
}

func (c CastBytesToInt64Err) Error() string {
	return fmt.Sprintf(
		"Cannot convert []byte to int64 because it contains less than 8 bytes (%d given)",
		c.BytesProvided,
	)
}

func BytesToInt64(bytes []byte, byteOrder binary.ByteOrder) (int64, error) {
	var bytesProvided int = len(bytes)
	if bytesProvided < 8 {
		return 0, CastBytesToInt64Err{BytesProvided: bytesProvided}
	}
	return int64(byteOrder.Uint64(bytes)), nil
}
