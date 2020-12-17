package fasttime

import (
	"strconv"
	"time"
)

// Strftime formats time t using format.
func Strftime(format string, t time.Time) string {
	return string(appendTime(make([]byte, 0, 64), format, t))
}

func appendTime(dst []byte, format string, t time.Time) []byte {
	n := len(format)
	if n == 0 {
		return dst
	}
	_ = format[n-1]
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	for i := 0; i < n; i++ {
		c := format[i]
		switch c {
		case '%':
			i++
			if i == n {
				return dst
			}
			switch format[i] {
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
					dst = strconv.AppendInt(dst, int64(hour%12), 10)
				case 'M':
					dst = strconv.AppendInt(dst, int64(minute), 10)
				case 'S':
					dst = strconv.AppendInt(dst, int64(second), 10)
				case 'j':
					dst = strconv.AppendInt(dst, int64(t.YearDay()), 10)
				}
			case 'a':
				dst = append(dst, t.Weekday().String()[:3]...)
			case 'A':
				dst = append(dst, t.Weekday().String()...)
			case 'b', 'h':
				dst = append(dst, month.String()[:3]...)
			case 'B':
				dst = append(dst, month.String()...)
			case 'c':
				dst = appendTime(dst, "%a %b %e %H:%M:%S %Y", t)
			case 'C':
				a := year / 100 * 2
				dst = append(dst, tab[a], tab[a+1])
			case 'd':
				dst = append(dst, tab[day*2], tab[day*2+1])
			case 'D':
				dst = appendTime(dst, "%m/%d/%y", t)
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
				dst = appendTime(dst, "%Y-%m-%d", t)
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
				a := hour % 12 * 2
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
			case 'r':
				dst = appendTime(dst, "%I:%M:%S %p", t)
			case 'R':
				dst = appendTime(dst, "%H:%M", t)
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
				dst = appendTime(dst, "%H:%M:%S", t)
			case 'u':
				dst = strconv.AppendInt(dst, int64(t.Weekday()+1), 10)
			case 'U':
				dst = strconv.AppendInt(dst, int64(((t.YearDay()-1)-int(t.Weekday()+6)%7+7)/7)+1, 10)
			case 'V':
				_, w := t.ISOWeek()
				dst = append(dst, tab[w*2], tab[w*2+1])
			case 'w':
				dst = strconv.AppendInt(dst, int64(t.Weekday()), 10)
			case 'W':
				dst = strconv.AppendInt(dst, int64(((t.YearDay()-1)-int(t.Weekday()+6)%7+7)/7), 10)
			case 'x':
				dst = appendTime(dst, dateFormat, t)
			case 'X':
				dst = appendTime(dst, timeFormat, t)
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
				case offset > 0:
					a := (offset / 60) / 60 * 2
					b := (offset / 60) % 60 * 2
					dst = append(dst, '+', tab[a], tab[a+1], tab[b], tab[b+1])
				}
			case 'Z':
				name, _ := t.Zone()
				dst = append(dst, name...)
			case '%':
				dst = append(dst, '%')
			}
		default:
			dst = append(dst, c)
		}
	}
	return dst
}
