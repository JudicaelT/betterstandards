package benchmark

import "testing"

func AssertAvgAllocs(
	b *testing.B,
	expectedAvgAllocs float64,
	codeUnderTest func(),
) {
	b.Run("allocs", func(b *testing.B) {
		var avgAllocs float64 = testing.AllocsPerRun(1000, codeUnderTest)
		if avgAllocs != expectedAvgAllocs {
			b.Fatalf(
				"Expected %.2f allocations on average, got %.2f instead",
				expectedAvgAllocs,
				avgAllocs,
			)
		}
	})
}
