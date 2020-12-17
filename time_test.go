package fasttime

import (
	"testing"
)

func TestTimestamp(t *testing.T) {
	t.Logf("%s\n", Timestamp())
}

func BenchmarkTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Timestamp()
	}
}
