package fasttime

import (
	"strconv"
	"time"
	"unsafe"
)

// Strftime formats time t using format.
func Strftime(format string, t time.Time) string {
	return string(AppendStrftime(make([]byte, 0, 64), format, t))
}

// AppendStrftime appends time t using format.
func AppendStrftime(dst []byte, format string, t time.Time) []byte {
	n := len(format)
	if n == 0 {
		return dst
	}
	_ = format[n-1]
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	for i := 0; i < n; i++ {
		c := format[i]
		if c != '%' {
			dst = append(dst, c)
			continue
		}
		i++
		if i == n {
			return dst
		}
		switch format[i] {
		case 'a':
			dst = append(dst, t.Weekday().String()[:3]...)
		case 'A':
			dst = append(dst, t.Weekday().String()...)
		case 'b', 'h':
			dst = append(dst, month.String()[:3]...)
		case 'B':
			dst = append(dst, month.String()...)
		case 'c':
			dst = AppendStrftime(dst, "%a %b %e %H:%M:%S %Y", t)
		case 'C':
			a := year / 100 * 2
			dst = append(dst, tab[a], tab[a+1])
		case 'd':
			dst = append(dst, tab[day*2], tab[day*2+1])
		case 'D':
			dst = AppendStrftime(dst, "%m/%d/%y", t)
		case 'e':
			if day >= 10 {
				dst = append(dst, tab[day*2], tab[day*2+1])
			} else {
				dst = append(dst, ' ', tab[day*2+1])
			}
		case 'E':
			panic("not implemented")
		case 'f':
			var tmp [6]byte
			a := t.Nanosecond() / 1000
			b := a % 100 * 2
			tmp[5] = tab[b+1]
			tmp[4] = tab[b]
			a /= 100
			b = a % 100 * 2
			tmp[3] = tab[b+1]
			tmp[2] = tab[b]
			b = a / 100 * 2
			tmp[1] = tab[b+1]
			tmp[0] = tab[b]
			dst = append(dst, tmp[:]...)
		case 'F':
			dst = AppendStrftime(dst, "%Y-%m-%d", t)
		case 'G':
			y, _ := t.ISOWeek()
			a := y / 100 * 2
			b := y % 100 * 2
			dst = append(dst, tab[a], tab[a+1], tab[b], tab[b+1])
		case 'g':
			y, _ := t.ISOWeek()
			a := y % 100 * 2
			dst = append(dst, tab[a], tab[a+1])
		case 'H':
			dst = append(dst, tab[hour*2], tab[hour*2+1])
		case 'I':
			a := hour % 12
			if a == 0 {
				a = 12
			}
			a *= 2
			dst = append(dst, tab[a], tab[a+1])
		case 'j':
			a := t.YearDay() / 100
			b := t.YearDay() % 100 * 2
			dst = append(dst, byte(a)+'0', tab[b], tab[b+1])
		case 'k':
			if hour >= 10 {
				dst = append(dst, tab[hour*2], tab[hour*2+1])
			} else {
				dst = append(dst, ' ', tab[hour*2+1])
			}
		case 'l':
			a := hour % 12
			if a == 0 {
				a = 12
			}
			if a >= 10 {
				dst = append(dst, tab[a*2], tab[a*2+1])
			} else {
				dst = append(dst, ' ', tab[a*2+1])
			}
		case 'm':
			dst = append(dst, tab[month*2], tab[month*2+1])
		case 'M':
			dst = append(dst, tab[minute*2], tab[minute*2+1])
		case 'n':
			dst = append(dst, '\n')
		case 'N':
			var tmp [9]byte
			a := t.Nanosecond()
			b := a % 100 * 2
			tmp[8] = tab[b+1]
			tmp[7] = tab[b]
			a /= 100
			b = a % 100 * 2
			tmp[6] = tab[b+1]
			tmp[5] = tab[b]
			a /= 100
			b = a % 100 * 2
			tmp[4] = tab[b+1]
			tmp[3] = tab[b]
			a /= 100
			b = a % 100 * 2
			tmp[2] = tab[b+1]
			tmp[1] = tab[b]
			tmp[0] = byte(a/100) + '0'
			dst = append(dst, tmp[:]...)
		case 'O':
			panic("not implemented")
		case 'p':
			if hour <= 12 {
				dst = append(dst, "AM"...)
			} else {
				dst = append(dst, "PM"...)
			}
		case 'P':
			if hour <= 12 {
				dst = append(dst, "am"...)
			} else {
				dst = append(dst, "pm"...)
			}
		case 'Q':
			a := t.Nanosecond() / 1000000
			b := a % 100 * 2
			dst = append(dst, byte(a/100)+'0', tab[b], tab[b+1])
		case 'r':
			dst = AppendStrftime(dst, "%I:%M:%S %p", t)
		case 'R':
			dst = AppendStrftime(dst, "%H:%M", t)
		case 's':
			var tmp [10]byte
			sec := t.Unix()
			is := sec % 100 * 2
			sec /= 100
			tmp[9] = tab[is+1]
			tmp[8] = tab[is]
			is = sec % 100 * 2
			sec /= 100
			tmp[7] = tab[is+1]
			tmp[6] = tab[is]
			is = sec % 100 * 2
			sec /= 100
			tmp[5] = tab[is+1]
			tmp[4] = tab[is]
			is = sec % 100 * 2
			sec /= 100
			tmp[3] = tab[is+1]
			tmp[2] = tab[is]
			is = sec % 100 * 2
			tmp[1] = tab[is+1]
			tmp[0] = tab[is]
			dst = append(dst, tmp[:]...)
		case 'S':
			dst = append(dst, tab[second*2], tab[second*2+1])
		case 't':
			dst = append(dst, '\t')
		case 'T':
			dst = AppendStrftime(dst, "%H:%M:%S", t)
		case 'u':
			dst = strconv.AppendInt(dst, int64(t.Weekday()+1), 10)
		case 'U':
			dst = strconv.AppendInt(dst, int64(((t.YearDay()-1)-int(t.Weekday()+6)%7+7)/7)+1, 10)
		case 'V':
			_, w := t.ISOWeek()
			dst = append(dst, tab[w*2], tab[w*2+1])
		case 'v':
			dst = AppendStrftime(dst, "%e-%b-%Y", t)
		case 'w':
			dst = strconv.AppendInt(dst, int64(t.Weekday()), 10)
		case 'W':
			dst = strconv.AppendInt(dst, int64(((t.YearDay()-1)-int(t.Weekday()+6)%7+7)/7), 10)
		case 'x':
			dst = AppendStrftime(dst, dateFormat, t)
		case 'X':
			dst = AppendStrftime(dst, timeFormat, t)
		case 'y':
			a := year % 100 * 2
			dst = append(dst, tab[a], tab[a+1])
		case 'Y':
			a := year / 100 * 2
			b := year % 100 * 2
			dst = append(dst, tab[a], tab[a+1], tab[b], tab[b+1])
		case 'z':
			_, offset := t.Zone()
			switch {
			case offset < 0:
				offset = -offset
				a := (offset / 60) / 60 * 2
				b := (offset / 60) % 60 * 2
				dst = append(dst, '-', tab[a], tab[a+1], tab[b], tab[b+1])
			case offset >= 0:
				a := (offset / 60) / 60 * 2
				b := (offset / 60) % 60 * 2
				dst = append(dst, '+', tab[a], tab[a+1], tab[b], tab[b+1])
			}
		case 'Z':
			name, _ := t.Zone()
			dst = append(dst, name...)
		case '%':
			dst = append(dst, '%')
		case '+':
			dst = AppendStrftime(dst, "%a %b %e %H:%M:%S %Z %Y", t)
		case '-':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
			case 'd':
				dst = strconv.AppendInt(dst, int64(day), 10)
			case 'm':
				dst = strconv.AppendInt(dst, int64(month), 10)
			case 'H':
				dst = strconv.AppendInt(dst, int64(hour), 10)
			case 'I':
				a := hour % 12
				if hour == 0 {
					a = 12
				}
				dst = strconv.AppendInt(dst, int64(a), 10)
			case 'M':
				dst = strconv.AppendInt(dst, int64(minute), 10)
			case 'S':
				dst = strconv.AppendInt(dst, int64(second), 10)
			case 'j':
				dst = strconv.AppendInt(dst, int64(t.YearDay()), 10)
			}
		case '_':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
			case 'd':
				if day < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(day), 10)
			case 'm':
				if month < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(month), 10)
			case 'H':
				if hour < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(hour), 10)
			case 'I':
				a := hour % 12
				if a == 0 {
					a = 12
				}
				if a < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(a), 10)
			case 'M':
				if minute < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(minute), 10)
			case 'S':
				if second < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(second), 10)
			case 'j':
				if second < 100 {
					dst = append(dst, ' ', ' ')
				} else if second < 10 {
					dst = append(dst, ' ')
				}
				dst = strconv.AppendInt(dst, int64(t.YearDay()), 10)
			}
		case '^':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
			case 'a':
				dst = appendUpper(dst, t.Weekday().String()[:3])
			case 'A':
				dst = appendUpper(dst, t.Weekday().String())
			case 'b', 'h':
				dst = appendUpper(dst, month.String()[:3])
			case 'B':
				dst = appendUpper(dst, month.String())
			case 'p':
				if hour <= 12 {
					dst = append(dst, "am"...)
				} else {
					dst = append(dst, "pm"...)
				}
			case 'P':
				if hour <= 12 {
					dst = append(dst, "AM"...)
				} else {
					dst = append(dst, "PM"...)
				}
			case 'r':
				dst = AppendStrftime(dst, "%I:%M:%S ", t)
				if hour <= 12 {
					dst = append(dst, "am"...)
				} else {
					dst = append(dst, "pm"...)
				}
			case 'x':
				b := AppendStrftime(nil, dateFormat, t)
				dst = appendUpper(dst, *(*string)(unsafe.Pointer(&b)))
			case 'X':
				b := AppendStrftime(nil, timeFormat, t)
				dst = appendUpper(dst, *(*string)(unsafe.Pointer(&b)))
			case 'Z':
				name, _ := t.Zone()
				dst = appendUpper(dst, name)
			}
		case '#':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
			case 'a':
				dst = appendSwapcase(dst, t.Weekday().String()[:3])
			case 'A':
				dst = appendSwapcase(dst, t.Weekday().String())
			case 'b', 'h':
				dst = appendSwapcase(dst, month.String()[:3])
			case 'B':
				dst = appendSwapcase(dst, month.String())
			case 'p':
				if hour <= 12 {
					dst = append(dst, "am"...)
				} else {
					dst = append(dst, "pm"...)
				}
			case 'P':
				if hour <= 12 {
					dst = append(dst, "AM"...)
				} else {
					dst = append(dst, "PM"...)
				}
			case 'r':
				dst = AppendStrftime(dst, "%I:%M:%S ", t)
				if hour <= 12 {
					dst = append(dst, "am"...)
				} else {
					dst = append(dst, "pm"...)
				}
			case 'x':
				b := AppendStrftime(nil, dateFormat, t)
				dst = appendSwapcase(dst, *(*string)(unsafe.Pointer(&b)))
			case 'X':
				b := AppendStrftime(nil, timeFormat, t)
				dst = appendSwapcase(dst, *(*string)(unsafe.Pointer(&b)))
			case 'Z':
				name, _ := t.Zone()
				dst = appendSwapcase(dst, name)
			}
		case ':':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
			case 'z':
				_, offset := t.Zone()
				switch {
				case offset == 0:
					dst = append(dst, 'Z')
				case offset < 0:
					offset = -offset
					a := (offset / 60) / 60 * 2
					b := (offset / 60) % 60 * 2
					dst = append(dst, '-', tab[a], tab[a+1], ':', tab[b], tab[b+1])
				case offset > 0:
					a := (offset / 60) / 60 * 2
					b := (offset / 60) % 60 * 2
					dst = append(dst, '+', tab[a], tab[a+1], ':', tab[b], tab[b+1])
				}
			}
		}
	}
	return dst
}

func appendUpper(dst []byte, s string) []byte {
	for _, c := range []byte(s) {
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		dst = append(dst, c)
	}
	return dst
}

func appendSwapcase(dst []byte, s string) []byte {
	for _, c := range []byte(s) {
		switch {
		case 'A' <= c && c <= 'Z':
			c += 'a' - 'A'
		case 'a' <= c && c <= 'z':
			c -= 'a' - 'A'
		}
		dst = append(dst, c)
	}
	return dst
}

const tab = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
