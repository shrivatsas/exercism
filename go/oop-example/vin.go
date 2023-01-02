package vin

// https://www.toptal.com/go/golang-oop-tutorial
import (
	"fmt"
)

// VIN is vehicle identification number
type VIN interface {
	Manufacturer() string
}

type vin string

// NewVIN is better for greping and leaves room for other
// NewXY funcs in the same package
func NewVIN(code string) (vin, error) {

	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	// ... check for disallowed characters ...

	return vin(code), nil
}

// Manufacturer is a function to determine the vehicle identification number
func (v vin) Manufacturer() string {

	return string(v[:3])
}

// vinEU is the format of european VINs
type vinEU vin

// NewEUVIN creates a EU VIN
func NewEUVIN(code string) (vinEU, error) {

	// call super constructor
	v, err := NewVIN(code)

	// and cast to subtype
	return vinEU(v), err
}

// Manufacturer returns only european ones
func (v vinEU) Manufacturer() string {

	// call manufacturer on supertype
	manufacturer := vin(v).Manufacturer()

	// add EU specific postfix if appropriate
	if manufacturer[2] == '9' {
		manufacturer += string(v[11:14])
	}

	return manufacturer
}
