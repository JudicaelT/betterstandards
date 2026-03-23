package cast

import (
	"encoding/binary"
	"fmt"
)

type CastBytesToInt32Err struct {
	BytesProvided int
}

func (c CastBytesToInt32Err) Error() string {
	return fmt.Sprintf(
		"Cannot convert []byte to int32 because it contains less than 4 bytes (%d given)",
		c.BytesProvided,
	)
}

func BytesToInt32(bytes []byte, byteOrder binary.ByteOrder) (int32, error) {
	var bytesProvided int = len(bytes)
	if bytesProvided < 4 {
		return 0, CastBytesToInt32Err{BytesProvided: bytesProvided}
	}
	return int32(byteOrder.Uint32(bytes)), nil
}
