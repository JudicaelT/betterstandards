package benchmark

import (
	"runtime"
	"testing"
)

func AssertNoAllocs(
	b *testing.B,
	codeUnderTest func(),
	cleanups ...func(),
) {
	b.Run("allocs", func(b *testing.B) {
		defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))

		var totalAllocs uint64 = 0
		var runs int = 1000

		for range runs {
			var before, after runtime.MemStats
			runtime.ReadMemStats(&before)
			codeUnderTest()
			runtime.ReadMemStats(&after)
			totalAllocs += after.Mallocs - before.Mallocs
			for _, cleanup := range cleanups {
				cleanup()
			}
		}

		avgAllocs := float64(totalAllocs / uint64(runs))
		if avgAllocs > 0 {
			b.Fatalf(
				"Expected zero allocations, got %.2f allocs on average",
				avgAllocs,
			)
		}
	})
}
