package benchmark

import (
	"runtime"
	"testing"
)

const avgAllocsRuns int = 1000

func AssertAvgAllocs(
	b *testing.B,
	expectedAvgAllocs float64,
	codeUnderTest func(),
	cleanups ...func(),
) {
	b.Run("allocs", func(b *testing.B) {
		defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))

		var totalAllocs uint64 = 0
		for range avgAllocsRuns {
			var before, after runtime.MemStats
			runtime.ReadMemStats(&before)
			codeUnderTest()
			runtime.ReadMemStats(&after)
			totalAllocs += after.Mallocs - before.Mallocs
			for _, cleanup := range cleanups {
				cleanup()
			}
		}

		avgAllocs := float64(totalAllocs / uint64(avgAllocsRuns))
		if avgAllocs != expectedAvgAllocs {
			b.Fatalf(
				"Expected %.2f allocations on average, got %.2f allocs on average",
				expectedAvgAllocs,
				avgAllocs,
			)
		}
	})
}
