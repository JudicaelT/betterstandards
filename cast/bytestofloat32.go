package cast

import (
	"encoding/binary"
	"fmt"
	"math"
)

type CastBytesToFloat32Err struct {
	BytesProvided int
}

func (c CastBytesToFloat32Err) Error() string {
	return fmt.Sprintf(
		"Cannot convert []byte to float32 because it contains less than 4 bytes (%d given)",
		c.BytesProvided,
	)
}

func BytesToFloat32(bytes []byte, byteOrder binary.ByteOrder) (float32, error) {
	var bytesProvided int = len(bytes)
	if bytesProvided < 4 {
		return 0, CastBytesToFloat32Err{BytesProvided: bytesProvided}
	}
	return math.Float32frombits(byteOrder.Uint32(bytes)), nil
}
