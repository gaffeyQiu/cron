package helper

import (
	"fmt"
)

func MustStringVar(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprint(v)
	case []byte:
		return string(v)
	case fmt.Stringer:
		return v.String()
	}
	return ""
}
