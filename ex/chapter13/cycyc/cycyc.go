package cycyc

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}

func cycyc(x reflect.Value, seen map[comparison]bool) bool {

	if !x.IsValid() {
		return false
	}

	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64,
		reflect.Complex128, reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return false

	case reflect.Ptr, reflect.Interface:
		return cycyc(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cycyc(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !cycyc(x.Field(i), seen) {
				return false
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if !cycyc(x.MapIndex(k), seen) {
				return false
			}
		}
		return false
	}
	panic("unreachable")
}

func Cycyc(x interface{}) bool {
	return cycyc(reflect.ValueOf(x), make(map[comparison]bool))
}
