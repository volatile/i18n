package i18n

import (
	"bytes"
	"strconv"

	"github.com/volatile/core"
)

// Num returns a formatted number with decimal and thousands marks, according to the locale decimalMark and thousandsMark respectively.
// If not set, the decimal mark is "," and the thousands mark is ".".
func Num(c *core.Context, n interface{}) string {
	switch n.(type) {
	case uint:
		return formatNum(c, []byte(strconv.FormatUint(uint64(n.(uint)), 10)))
	case uint8:
		return formatNum(c, []byte(strconv.FormatUint(uint64(n.(uint8)), 10)))
	case uint16:
		return formatNum(c, []byte(strconv.FormatUint(uint64(n.(uint16)), 10)))
	case uint32:
		return formatNum(c, []byte(strconv.FormatUint(uint64(n.(uint32)), 10)))
	case uint64:
		return formatNum(c, []byte(strconv.FormatUint(n.(uint64), 10)))
	case int:
		return formatNum(c, []byte(strconv.Itoa(n.(int))))
	case int8:
		return formatNum(c, []byte(strconv.FormatInt(int64(n.(int8)), 10)))
	case int16:
		return formatNum(c, []byte(strconv.FormatInt(int64(n.(int16)), 10)))
	case int32:
		return formatNum(c, []byte(strconv.FormatInt(int64(n.(int32)), 10)))
	case int64:
		return formatNum(c, []byte(strconv.FormatInt(n.(int64), 10)))
	case float32:
		println(strconv.FormatFloat(float64(n.(float32)), 'f', 8, 32))
		return formatNum(c, []byte(strconv.FormatFloat(float64(n.(float32)), 'f', -1, 32)))
	case float64:
		return formatNum(c, []byte(strconv.FormatFloat(n.(float64), 'f', -1, 64)))
	case string:
		return formatNum(c, []byte(n.(string)))
	case []byte:
		return formatNum(c, n.([]byte))
	default:
		return ""
	}
}

func formatNum(c *core.Context, b []byte) (s string) {
	decimalMark := "."
	thousandsMark := ","
	if locale, ok := locales[ClientLocale(c)]; ok {
		if v, ok := locale["decimalMark"]; ok {
			decimalMark = v
		}
		if v, ok := locale["thousandsMark"]; ok {
			thousandsMark = v
		}
	}

	bb := bytes.Split(b, []byte("."))
	switch len(bb) {
	case 1:
		break
	case 2:
		s = decimalMark + string(bb[1])
		b = bb[0]
	default:
		return // Can't have 2 decimal marks in a number so return nothing.
	}

	j := 0
	for i := len(b) - 1; i >= 0; i-- {
		if j != 0 && j%3 == 0 {
			s = thousandsMark + s
		}
		s = string(b[i]) + s
		j++
	}
	return
}
