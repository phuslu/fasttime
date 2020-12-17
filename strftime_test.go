package fasttime

import (
	"testing"
	"time"
)

var now = time.Date(2006, 1, 2, 15, 03, 04, 0, time.FixedZone("MST", -25200))

func TestStrftime(t *testing.T) {
	cases := []struct {
		Format string
		Result string
	}{
		{"%a", "Mon"},
		{"%A", "Monday"},
		{"%b", "Jan"},
		{"%B", "January"},
		{"%c", "Mon Jan  2 15:03:04 2006"},
		{"%C", "20"},
		{"%d", "02"},
		{"%D", "01/02/06"},
		{"%e", " 2"},
		{"%F", "2006-01-02"},
		{"%G", "2006"},
		{"%g", "06"},
		{"%h", "Jan"},
		{"%H", "15"},
		{"%I", "03"},
		{"%j", "002"},
		{"%k", "15"},
		{"%l", " 3"},
		{"%m", "01"},
		{"%M", "03"},
		{"%n", "\n"},
		{"%p", "PM"},
		{"%P", "pm"},
		{"%r", "03:03:04 PM"},
		{"%R", "15:03"},
		{"%s", "1136239384"},
		{"%S", "04"},
		{"%t", "\t"},
		{"%T", "15:03:04"},
		// {"%u", "1"},
		// {"%U", "01"},
		{"%V", "01"},
		{"%w", "1"},
		// {"%W", "01"},
		// {"%x", "02/01/2006"},
		// {"%X", "15:03:04"},
		{"%y", "06"},
		{"%Y", "2006"},
		{"%z", "-0700"},
		{"%Z", "MST"},
		{"%%", "%"},
	}

	for _, c := range cases {
		if got, want := Strftime(c.Format, now), c.Result; got != want {
			t.Errorf("Strftime(%+v, %v) want=%v got=%v", c.Format, now, want, got)
		}
	}
}

func BenchmarkISOTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strftime("%Y-%m-%d %H:%M:%S", now)
	}
}
