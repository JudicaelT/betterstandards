package benchmark

import "testing"

func AvgRuntime(
	b *testing.B,
	codeUnderTest func(),
	cleanups ...func(),
) {
	b.Run("runtime", func(b *testing.B) {
		for b.Loop() {
			codeUnderTest()
			b.StopTimer()
			for _, cleanup := range cleanups {
				cleanup()
			}
			b.StartTimer()
		}
	})
}
