package fasttime

// Timestamp formats unix timestamp.
func Timestamp() string {
	return string(timestamp(make([]byte, 10)))
}

func timestamp(b []byte) []byte {
	sec, _ := walltime()
	a := sec % 100 * 2
	sec /= 100
	b[9] = tab[a+1]
	b[8] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[7] = tab[a+1]
	b[6] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[5] = tab[a+1]
	b[4] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[3] = tab[a+1]
	b[2] = tab[a]
	a = sec % 100 * 2
	b[1] = tab[a+1]
	b[0] = tab[a]
	return b
}

// TimestampMS formats unix timestamp with milli seconds.
func TimestampMS() string {
	return string(timestampms(make([]byte, 13)))
}

func timestampms(b []byte) []byte {
	sec, nsec := walltime()
	ms := int64(nsec) / 1000000
	a := ms % 100 * 2
	b[12] = tab[a+1]
	b[11] = tab[a]
	b[10] = byte('0' + ms/100)
	a = sec % 100 * 2
	sec /= 100
	b[9] = tab[a+1]
	b[8] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[7] = tab[a+1]
	b[6] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[5] = tab[a+1]
	b[4] = tab[a]
	a = sec % 100 * 2
	sec /= 100
	b[3] = tab[a+1]
	b[2] = tab[a]
	a = sec % 100 * 2
	b[1] = tab[a+1]
	b[0] = tab[a]
	return b
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
