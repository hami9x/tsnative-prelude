package tsn

import (
	"fmt"
	"reflect"
	"strconv"
)

var Tsn__ = tsnUtils{}

type tsnUtils struct{}

func (tsnUtils) formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// number to string
func (u tsnUtils) ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case float32:
		return u.formatFloat(float64(v))
	case float64:
		return u.formatFloat(v)
	default:
		if u.interfaceHasNilValue(value) {
			return ""
		}
		return fmt.Sprintf("%+v", v)
	}
	panic("unreachable")
}

func (u tsnUtils) interfaceHasNilValue(actual interface{}) bool {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	nilable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map

	// Careful: reflect.Value.IsNil() will panic unless it's an interface, chan, map, func, slice, or ptr
	// Reference: http://golang.org/pkg/reflect/#Value.IsNil
	return nilable && value.IsNil()
}
