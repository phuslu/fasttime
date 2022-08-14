# fasttime - fast time formatting for go

[![godoc][godoc-img]][godoc] [![release][release-img]][release] [![goreport][goreport-img]][goreport] [![coverage][coverage-img]][coverage]

## Getting Started

try on https://go.dev/play/p/wnb6G181qu6
```go
package main

import (
	"time"

	"github.com/phuslu/fasttime"
)

func main() {
	println(fasttime.Strftime("%a %b %d %H:%M:%S %Z %Y", time.Now()))
}

// Thu Dec 17 10:49:04 +08 2020
```

## Benchmarks

```go
// go test -v -cpu=4 -run=none -bench=. -benchmem strftime_test.go
package strftime_test

import (
	"io"
	"testing"
	"time"

	itchyny "github.com/itchyny/timefmt-go"
	lestrrat "github.com/lestrrat-go/strftime"
	fasttime "github.com/phuslu/fasttime"
)

var now = time.Now()

func BenchmarkUnixDateStdTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now.Format(time.UnixDate)
	}
}

func BenchmarkUnixDateLestrrat(b *testing.B) {
	p, _ := lestrrat.New("%a %b %e %H:%M:%S %Z %Y")
	for i := 0; i < b.N; i++ {
		p.Format(io.Discard, now)
	}
}

func BenchmarkUnixDateItchyny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		itchyny.Format(now, "%a %b %e %H:%M:%S %Z %Y")
	}
}

func BenchmarkUnixDateFasttime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fasttime.Strftime("%a %b %e %H:%M:%S %Z %Y", now)
	}
}

func BenchmarkStampMicroStdTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now.Format(time.StampMicro)
	}
}

func BenchmarkStampMicroLestrrat(b *testing.B) {
	p, _ := lestrrat.New("%b %e %H:%M:%S.%f", lestrrat.WithMicroseconds('f'))
	for i := 0; i < b.N; i++ {
		p.Format(io.Discard, now)
	}
}

func BenchmarkStampMicroItchyny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		itchyny.Format(now, "%b %e %H:%M:%S.%f")
	}
}

func BenchmarkStampMicroFasttime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fasttime.Strftime("%b %e %H:%M:%S.%f", now)
	}
}

func BenchmarkRFC3339StdTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now.Format(time.RFC3339)
	}
}

func BenchmarkRFC3339Lestrrat(b *testing.B) {
	p, _ := lestrrat.New("%Y-%m-%dT%H:%M:%S%z")
	for i := 0; i < b.N; i++ {
		p.Format(io.Discard, now)
	}
}

func BenchmarkRFC3339Itchyny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		itchyny.Format(now, "%Y-%m-%dT%H:%M:%S%:z")
	}
}

func BenchmarkRFC3339Fasttime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fasttime.Strftime("%Y-%m-%dT%H:%M:%S%:z", now)
	}
}
```
A Performance result as below, for daily benchmark results see [github actions][benchmark]
```
BenchmarkUnixDateStdTime-4      	 3134547	       359.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkUnixDateLestrrat-4     	 2445044	       505.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnixDateItchyny-4      	 3726554	       316.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnixDateFasttime-4     	 4719768	       245.8 ns/op	       0 B/op	       0 allocs/op

BenchmarkStampMicroStdTime-4    	 3793047	       318.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkStampMicroLestrrat-4   	 2409372	       477.2 ns/op	      72 B/op	       2 allocs/op
BenchmarkStampMicroItchyny-4    	 4752626	       242.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkStampMicroFasttime-4   	 7268750	       164.2 ns/op	       0 B/op	       0 allocs/op

BenchmarkRFC3339StdTime-4       	 3563881	       336.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkRFC3339Lestrrat-4      	 2520579	       480.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkRFC3339Itchyny-4       	 4386028	       273.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkRFC3339Fasttime-4      	 6393230	       189.1 ns/op	       0 B/op	       0 allocs/op
```

## Supported formats:

| code | Description |
| ---- | --- |
| `%a` | The abbreviated name of the day of the week according to the current locale. |
| `%A` | The full name of the day of the week according to the current locale. |
| `%b` | The abbreviated month name according to the current locale. |
| `%B` | The full month name according to the current locale. |
| `%c` | The preferred date and time representation for the current locale. |
| `%C` | The century number (year/100) as a 2-digit integer. |
| `%d` | The day of the month as a decimal number (range 01 to 31). |
| `%D` | Equivalent to %m/%d/%y.  (Yecch—for Americans only.) |
| `%e` | Like %d, the day of the month as a decimal number, but a leading zero is replaced by a space. |
| `%F` | Equivalent to %Y-%m-%d (the ISO 8601 date format). |
| `%G` | The  ISO 8601 week-based year (see NOTES) with century as a decimal number.  The 4-digit year corresponding to the ISO week number (see %V).  This has the same format and value as %Y, except that if the ISO week number belongs to the previous or next year, that year is used instead. |
| `%g` | Like %G, but without century, that is, with a 2-digit year (00-99). |
| `%h` | Equivalent to %b. |
| `%H` | The hour as a decimal number using a 24-hour clock (range 00 to 23). |
| `%I` | The hour as a decimal number using a 12-hour clock (range 01 to 12). |
| `%j` | The day of the year as a decimal number (range 001 to 366). |
| `%k` | The hour (24-hour clock) as a decimal number (range 0 to 23); single digits are preceded by a blank.  (See also %H.)  |
| `%l` | The hour (12-hour clock) as a decimal number (range 1 to 12); single digits are preceded by a blank.  (See also %I.)  |
| `%m` | The month as a decimal number (range 01 to 12). |
| `%M` | The minute as a decimal number (range 00 to 59). |
| `%n` | A newline character. |
| `%p` | Either "AM" or "PM" according to the given time value, or the corresponding strings for the current locale.  Noon is treated as "PM" and midnight as "AM". |
| `%P` | Like %p but in lowercase: "am" or "pm" or a corresponding string for the current locale. |
| `%r` | The time in a.m. or p.m. notation.  In the POSIX locale this is equivalent to %I:%M:%S %p. |
| `%R` | The time in 24-hour notation (%H:%M). For a version including the seconds, see %T below. |
| `%s` | The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC). |
| `%S` | The second as a decimal number (range 00 to 60).  (The range is up to 60 to allow for occasional leap seconds.) |
| `%t` | A tab character. |
| `%T` | The time in 24-hour notation (%H:%M:%S). |
| `%u` | The day of the week as a decimal, range 1 to 7, Monday being 1.  See also %w. |
| `%U` | The week number of the current year as a decimal number, range 00 to 53, starting with the first Sunday as the first day of week 01.  See also %V and %W. |
| `%v` | Equivalent to %e-%b-%Y. |
| `%V` | The ISO 8601 week number (see NOTES) of the current year as a decimal number, range 01 to 53, where week 1 is the first week that has at least 4 days in the new year.  See also %U and %W. |
| `%w` | The day of the week as a decimal, range 0 to 6, Sunday being 0.  See also %u. |
| `%W` | The week number of the current year as a decimal number, range 00 to 53, starting with the first Monday as the first day of week 01. |
| `%x` | The preferred date representation for the current locale without the time. |
| `%X` | The preferred time representation for the current locale without the date. |
| `%y` | The year as a decimal number without a century (range 00 to 99). |
| `%Y` | The year as a decimal number including the century. |
| `%z` | The +hhmm or -hhmm numeric timezone (that is, the hour and minute offset from UTC). |
| `%Z` | The timezone name or abbreviation. |
| `%+` | The date and time in date(1) format. |
| `%%` | A literal '%' character. |

[godoc-img]: http://img.shields.io/badge/godoc-reference-blue.svg
[godoc]: https://godoc.org/github.com/phuslu/fasttime
[release-img]: https://img.shields.io/github/v/tag/phuslu/fasttime?label=release
[release]: https://github.com/phuslu/fasttime/releases
[goreport-img]: https://goreportcard.com/badge/github.com/phuslu/fasttime
[goreport]: https://goreportcard.com/report/github.com/phuslu/fasttime
[coverage-img]: http://gocover.io/_badge/github.com/phuslu/fasttime
[coverage]: https://gocover.io/github.com/phuslu/fasttime
[benchmark]: https://github.com/phuslu/fasttime/actions/workflows/benchmark.yml?query=workflow%3Abenchmark
