package fasttime

const (
	Layout      = "%m/%d %I:%M:%S%p '%y %z"  // 01/02 03:04:05PM '06 -0700
	ANSIC       = "%a %b %e %H:%M:%S %Y"     // Mon Jan _2 15:04:05 2006
	UnixDate    = "%a %b %e %H:%M:%S %Z %Y"  // Mon Jan _2 15:04:05 MST 2006
	RubyDate    = "%a %b %d %H:%M:%S %z %Y"  // Mon Jan 02 15:04:05 -0700 2006
	RFC822      = "%d %b %y %H:%M %Z"        // 02 Jan 06 15:04 MST
	RFC822Z     = "%d %b %y %H:%M %z"        // 02 Jan 06 15:04 -0700
	RFC850      = "%A, %d-%b-%y %H:%M:%S %Z" // Monday, 02-Jan-06 15:04:05 MST
	RFC1123     = "%a, %d %b %Y %H:%M:%S %Z" // Mon, 02 Jan 2006 15:04:05 MST
	RFC1123Z    = "%a, %d %b %Y %H:%M:%S %z" // Mon, 02 Jan 2006 15:04:05 -0700
	RFC3339     = "%Y-%m-%dT%H:%M:%S%:z"     // 2006-01-02T15:04:05Z07:00
	RFC3339Nano = "%Y-%m-%dT%H:%M:%S.%N%:z"  // 2006-01-02T15:04:05.000000000Z07:00
	Kitchen     = "%-I:%M%p"                 // 3:04PM
	// Handy time stamps.
	Stamp      = "%b %e %H:%M:%S"    // Jan _2 15:04:05
	StampMilli = "%b %e %H:%M:%S.%Q" // Jan _2 15:04:05.000
	StampMicro = "%b %e %H:%M:%S.%f" // Jan _2 15:04:05.000000
	StampNano  = "%b %e %H:%M:%S.%N" // Jan _2 15:04:05.000000000
)
