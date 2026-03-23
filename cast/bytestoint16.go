package cast

import (
	"encoding/binary"
	"fmt"
)

type CastBytesToInt16Err struct {
	BytesProvided int
}

func (c CastBytesToInt16Err) Error() string {
	return fmt.Sprintf(
		"Cannot convert []byte to int16 because it contains less than 2 bytes (%d given)",
		c.BytesProvided,
	)
}

func BytesToInt16(bytes []byte, byteOrder binary.ByteOrder) (int16, error) {
	var bytesProvided int = len(bytes)
	if bytesProvided < 2 {
		return 0, CastBytesToInt16Err{BytesProvided: bytesProvided}
	}
	return int16(byteOrder.Uint16(bytes)), nil
}
