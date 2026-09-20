package fasttime

import (
	"testing"
	"time"
)

func TestLayout(t *testing.T) {
	cases := []struct {
		Layout    string
		StdLayout string
	}{
		{Layout, time.Layout},
		{ANSIC, time.ANSIC},
		{UnixDate, time.UnixDate},
		{RubyDate, time.RubyDate},
		{RFC822, time.RFC822},
		{RFC822Z, time.RFC822Z},
		{RFC850, time.RFC850},
		{RFC1123, time.RFC1123},
		{RFC1123Z, time.RFC1123Z},
		{RFC3339, time.RFC3339},
		{RFC3339Milli, "2006-01-02T15:04:05.999Z07:00"},
		{RFC3339Micro, "2006-01-02T15:04:05.999999Z07:00"},
		{RFC3339Nano, time.RFC3339Nano},
		{Kitchen, time.Kitchen},
		{Stamp, time.Stamp},
		{StampMilli, time.StampMilli},
		{StampMicro, time.StampMicro},
		{StampNano, time.StampNano},
	}

	now := time.Now()

	for _, c := range cases {
		for _, rounding := range []time.Duration{time.Nanosecond, time.Nanosecond * 10, time.Nanosecond * 100, time.Microsecond, time.Microsecond * 10, time.Microsecond * 100, time.Millisecond, time.Millisecond * 10, time.Millisecond * 100, time.Second, time.Minute} {
			now := now.Round(rounding)
			if got, want := Strftime(c.Layout, now), now.Format(c.StdLayout); got != want {
				t.Errorf("Strftime(%#v, %#v) want=%v got=%v", c.Layout, atime, want, got)
			}
		}
	}
}
