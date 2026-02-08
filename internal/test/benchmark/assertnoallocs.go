package benchmark

import "testing"

func AssertNoAllocs(
	b *testing.B,
	codeUnderTest func(),
) {
	b.Run("allocs", func(b *testing.B) {
		avgAllocs := testing.AllocsPerRun(1000, codeUnderTest)
		if avgAllocs != 0 {
			b.Fatalf(
				"Expected zero allocations, got %.2f allocs on average",
				avgAllocs,
			)
		}
	})
}
