package comcom

import (
	"testing"
)

func TestToFloat64(t *testing.T) {

	if x, err := ToFloat64(true); err != nil {
		t.Errorf("ToFloat64(true): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(true) != 1.0")
	}

	if x, err := ToFloat64(int(1)); err != nil {
		t.Errorf("ToFloat64(int(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(int(1)) != 1.0")
	}

	if x, err := ToFloat64(int8(1)); err != nil {
		t.Errorf("ToFloat64(int8(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(int8(1)) != 1.0")
	}

	if x, err := ToFloat64(int16(1)); err != nil {
		t.Errorf("ToFloat64(int16(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(int16(1)) != 1.0")
	}

	if x, err := ToFloat64(int32(1)); err != nil {
		t.Errorf("ToFloat64(int32(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(int32(1)) != 1.0")
	}

	if x, err := ToFloat64(int64(1)); err != nil {
		t.Errorf("ToFloat64(int64(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(int64(1)) != 1.0")
	}

	if x, err := ToFloat64(uint(1)); err != nil {
		t.Errorf("ToFloat64(uint(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uint(1)) != 1.0")
	}

	if x, err := ToFloat64(uint8(1)); err != nil {
		t.Errorf("ToFloat64(uint8(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uint8(1)) != 1.0")
	}

	if x, err := ToFloat64(uint16(1)); err != nil {
		t.Errorf("ToFloat64(uint16(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uint16(1)) != 1.0")
	}

	if x, err := ToFloat64(uint32(1)); err != nil {
		t.Errorf("ToFloat64(uint32(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uint32(1)) != 1.0")
	}

	if x, err := ToFloat64(uint64(1)); err != nil {
		t.Errorf("ToFloat64(uint64(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uint64(1)) != 1.0")
	}

	if x, err := ToFloat64(uintptr(1)); err != nil {
		t.Errorf("ToFloat64(uintptr(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(uintptr(1)) != 1.0")
	}

	if x, err := ToFloat64(float32(1)); err != nil {
		t.Errorf("ToFloat64(float32(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(float32(1)) != 1.0")
	}

	if x, err := ToFloat64(float64(1)); err != nil {
		t.Errorf("ToFloat64(float64(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(float64(1)) != 1.0")
	}

	if x, err := ToFloat64(complex64(1)); err != nil {
		t.Errorf("ToFloat64(complex64(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(complex64(1)) != 1.0")
	}

	if x, err := ToFloat64(complex128(1)); err != nil {
		t.Errorf("ToFloat64(complex128(1)): %v", err)
	} else if x != 1.0 {
		t.Error("ToFloat64(complex128(1)) != 1.0")
	}

	if _, err := ToFloat64(complex64(1 + 1i)); err == nil {
		t.Error("ToFloat64(complex64(1+1i)) didn't report an error")
	}

	if _, err := ToFloat64(complex128(1 + 1i)); err == nil {
		t.Error("ToFloat64(complex128(1+1i)) didn't report an error")
	}

}

func TestComcom(t *testing.T) {

	for _, test := range []struct {
		x interface{}
		y interface{}
		want bool
	}{
		{1, 1.0, true},
		{true, false, false},
		{uint(1), 1+0i, true},
		{-1, -1-0i, true},
	}{
		if got, _ := Comcom(test.x, test.y); got != test.want {
			t.Errorf("Comcom(%v,%v) = %t", test.x, test.y, got)
		}
	}

}
