// +build windows

package fasttime

import (
	"strings"
	"syscall"
	"unsafe"
)

var dateFormat, timeFormat = func() (string, string) {
	var h syscall.Handle
	var b [64]uint16
	var n uint32

	replacer := strings.NewReplacer(
		"yyyy", "%Y",
		"yy", "%y",
		"MMMM", "%B",
		"MM", "%m",
		"M", "%-m",
		"dddd", "%A",
		"dd", "%d",
		"d", "%-d",
		"HH", "%H",
		"H", "%-H",
		"hh", "%I",
		"h", "%-I",
		"mm", "%M",
		"ss", "%S",
		"tt", "%P",
	)

	// open registry
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(`Control Panel\International`), 0, syscall.KEY_READ, &h)
	if err != nil {
		return "", ""
	}
	defer syscall.RegCloseKey(h)

	// read sShortDate
	n = uint32(len(b))
	err = syscall.RegQueryValueEx(h, syscall.StringToUTF16Ptr(`sShortDate`), nil, nil, (*byte)(unsafe.Pointer(&b[0])), &n)
	if err != nil {
		return "", ""
	}
	d := replacer.Replace(syscall.UTF16ToString(b[:n]))

	// read sTimeFormat
	n = uint32(len(b))
	err = syscall.RegQueryValueEx(h, syscall.StringToUTF16Ptr(`sTimeFormat`), nil, nil, (*byte)(unsafe.Pointer(&b[0])), &n)
	if err != nil {
		return "", ""
	}
	t := replacer.Replace(syscall.UTF16ToString(b[:n]))

	return d, t
}()
