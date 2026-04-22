package benchmark

import "testing"

func AvgRuntime(
	b *testing.B,
	codeUnderTest func(),
	cleanups ...func(),
) {
	b.Run("runtime", func(b *testing.B) {
		// Adding cleanups will cause the timer to stop which triggers the slow path.
		// Therefore, we check if there are no cleanup functions to avoid the slow path.
		if len(cleanups) == 0 {
			for b.Loop() {
				codeUnderTest()
			}
		} else {
			codeUnderTest()
			b.StopTimer()
			for _, cleanup := range cleanups {
				cleanup()
			}
			b.StartTimer()
		}
	})
}
