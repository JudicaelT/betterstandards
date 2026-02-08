package benchmark

import (
	"testing"
	"time"
)

func AssertMaxDuration(
	b *testing.B,
	codeUnderTest func(),
	maxExecutionDuration time.Duration,
) {
	b.Run("duration", func(b *testing.B) {
		for b.Loop() {
			codeUnderTest()
		}
		var avgNanoseconds int64 = b.Elapsed().Nanoseconds() / int64(b.N)
		if avgNanoseconds > maxExecutionDuration.Nanoseconds() {
			b.Errorf(
				"Expected code to run for %dns or less. Ran for %dns on average",
				maxExecutionDuration.Nanoseconds(),
				avgNanoseconds,
			)
		}
	})
}
