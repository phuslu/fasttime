package fasttime

import (
	"strconv"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	t.Logf("%s\n", Timestamp())
}

func BenchmarkStdTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(time.Now().Unix(), 10)
	}
}

func BenchmarkFastTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Timestamp()
	}
}
