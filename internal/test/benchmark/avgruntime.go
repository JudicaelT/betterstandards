package benchmark

import "testing"

func AvgRuntime(
	b *testing.B,
	codeUnderTest func(),
) {
	b.Run("runtime", func(b *testing.B) {
		for b.Loop() {
			codeUnderTest()
		}
	})
}
