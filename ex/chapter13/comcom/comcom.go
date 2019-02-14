package comcom

import (
	"fmt"
	"reflect"
)

func toFloat64(x reflect.Value) (X float64, err error) {
	if !x.IsValid() {
		return 0.0, fmt.Errorf("argument is reflect-invalid")
	}

	switch x.Kind() {
	case reflect.String, reflect.Chan, reflect.UnsafePointer, reflect.Func, reflect.Ptr,
		reflect.Interface, reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
		return 0.0, fmt.Errorf("expected a number, got %v", x.Type())
	case reflect.Bool:
		if x.Interface().(bool) {
			X = 1.0
		}
	case reflect.Int:
		X = float64(x.Interface().(int))
	case reflect.Int8:
		X = float64(x.Interface().(int8))
	case reflect.Int16:
		X = float64(x.Interface().(int16))
	case reflect.Int32:
		X = float64(x.Interface().(int32))
	case reflect.Int64:
		X = float64(x.Interface().(int64))
	case reflect.Uint:
		X = float64(x.Interface().(uint))
	case reflect.Uint8:
		X = float64(x.Interface().(uint8))
	case reflect.Uint16:
		X = float64(x.Interface().(uint16))
	case reflect.Uint32:
		X = float64(x.Interface().(uint32))
	case reflect.Uint64:
		X = float64(x.Interface().(uint64))
	case reflect.Uintptr:
		X = float64(x.Interface().(uintptr))
	case reflect.Float32:
		X = float64(x.Interface().(float32))
	case reflect.Float64:
		X = float64(x.Interface().(float64))
	case reflect.Complex64:
		z := x.Interface().(complex64)
		if imag(z) != 0 {
			return 0.0, fmt.Errorf("uncomparable complex number: %v", z)
		} else {
			X = float64(real(z))
		}
	case reflect.Complex128:
		z := x.Interface().(complex128)
		if imag(z) != 0 {
			return 0.0, fmt.Errorf("uncomparable complex number: %v", z)
		} else {
			X = float64(real(z))
		}
	}

	return X, nil
}

func abs(x float64) float64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func comcom(x reflect.Value, y reflect.Value) (bool, error) {
	X, err := toFloat64(x)
	if err != nil {
		return false, fmt.Errorf("first argument is bad: %v", err)
	}
	Y, err := toFloat64(y)
	if err != nil {
		return false, fmt.Errorf("first argument is bad: %v", err)
	}

	const EPS = 0.00001
	return abs(X-Y) < EPS, nil
}

func ToFloat64(x interface{}) (float64, error) {
	return toFloat64(reflect.ValueOf(x))
}

func Comcom(x interface{}, y interface{}) (bool, error) {
	return comcom(reflect.ValueOf(x), reflect.ValueOf(y))
}
