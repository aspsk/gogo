package tempconv

import (
        "testing"
)

func TestC2K(t *testing.T) {
	if C2K(AbsoluteZero) != 0 {
		t.Error("C2K(AbsoluteZero)=", C2K(AbsoluteZero))
	}
}

func TestK2C(t *testing.T) {
	if K2C(0) != AbsoluteZero {
		t.Error("K2C(0)=", K2C(0))
	}
}

func TestC2F(t *testing.T) {
	if x := C2F(0); x != 32 {
		t.Error("C2F(0)=", x)
	}

	if x := C2F(100); x != 212 {
		t.Error("C2F(100)=", x)
	}
}

func TestF2C(t *testing.T) {
	if x := F2C(32); x != 0 {
		t.Error("F2C(0)=", x)
	}

	if x := F2C(212); x != 100 {
		t.Error("F2C(212)=", x)
	}
}
