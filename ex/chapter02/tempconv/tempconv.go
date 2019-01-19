package tempconv

import "fmt"

type basetype float64

type Celsius basetype
type Kelvin basetype
type Fahrenheit basetype

const (

    AbsoluteZero Celsius = -273.15
    Freezing Celsius = 0
    Boiling Celsius = 100

)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func C2K(c Celsius) Kelvin {
    return Kelvin(basetype(c - AbsoluteZero))
}

func C2F(c Celsius) Fahrenheit {
    return Fahrenheit(basetype(c * 9 / 5.0 + 32))
}

func K2C(k Kelvin) Celsius {
    return Celsius(basetype(k) + basetype(AbsoluteZero))
}

func K2F(k Kelvin) Fahrenheit {
    return C2F(K2C(k))
}

func F2C(f Fahrenheit) Celsius {
    return Celsius(basetype(f - 32) * 5 / 9.0)
}

func F2K(f Fahrenheit) Kelvin {
    return C2K(F2C(f))
}
