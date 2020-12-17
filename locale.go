// +build !windows

package fasttime

import (
	"os"
	"strings"
)

var dateFormat, timeFormat = func() (string, string) {
	lang := os.Getenv("LC_TIME")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if i := strings.IndexByte(lang, '.'); i > 0 {
		lang = lang[:i]
	}
	if lang == "" {
		lang = os.Getenv("LANG")
	}

	switch lang {
	case "en_AG":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_AU":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_BW":
		return "%d/%m/%Y", "%H:%M:%S"
	case "en_CA":
		return "%Y-%m-%d", "%I:%M:%S %p"
	case "en_DK":
		return "%Y-%m-%d", "%H:%M:%S"
	case "en_GB":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_HK":
		return "%A, %B %d, %Y", "%I:%M:%S %Z"
	case "en_IE":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_IL":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_IN":
		return "%A %d %B %Y", "%I:%M:%S  %Z"
	case "en_NG":
		return "%d/%m/%Y", "%H:%M:%S"
	case "en_NZ":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_PH":
		return "%A, %d %B, %Y", "%I:%M:%S  %Z"
	case "en_SC":
		return "%m/%d/%y", "%H:%M:%S"
	case "en_SG":
		return "%d/%m/%Y", "%H:%M:%S"
	case "en_US":
		return "%m/%d/%Y", "%I:%M:%S %p"
	case "en_ZA":
		return "%d/%m/%Y", "%H:%M:%S"
	case "en_ZM":
		return "%d/%m/%y", "%H:%M:%S"
	case "en_ZW":
		return "%d/%m/%Y", "%H:%M:%S"
	default:
		return "%m/%d/%Y", "%H:%M:%S"
	}
}()
