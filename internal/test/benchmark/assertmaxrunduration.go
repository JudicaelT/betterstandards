package benchmark

import (
	"testing"
	"time"
)

func AssertMaxRunDuration(
	b *testing.B,
	functionUnderTest string,
	maxExecutionDuration time.Duration,
) {
	recover()
	b.StopTimer()
	if b.Elapsed() > maxExecutionDuration {
		b.Errorf(
			"%s() ran for %dns but should have taken %dns or less",
			functionUnderTest,
			b.Elapsed().Nanoseconds(),
			maxExecutionDuration.Nanoseconds(),
		)
	}
}
