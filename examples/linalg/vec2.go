package main

import (
	"fmt"
	"math/cmplx"
)

type vec2 complex128

func (v vec2) x() float64         { return real(v) }
func (v vec2) y() float64         { return imag(v) }
func (v vec2) String() string     { return fmt.Sprintf("%.2f %.2f", v.x(), v.y()) }
func (v vec2) Magnitude() float64 { return cmplx.Abs(complex128(v)) }

func (v vec2) Normalize() vec2 {
	mag := v.Magnitude()
	if mag == 0 {
		return vec2(0)
	}
	return v * newVec2(1/mag, 0)
}

func Dot(a, b vec2) float64 {
	return a.x()*b.x() + a.y()*b.y()
}

func newVec2(x, y float64) vec2 {
	return vec2(complex(x, y))
}

func main() {
	pos := newVec2(2, 2)
	vel := newVec2(1, 0)

	fmt.Println(pos)
	fmt.Println(vel)
	fmt.Println(pos + vel)

	fmt.Println(pos.x(), pos.y())

	fmt.Println(pos.Magnitude())
	fmt.Println(pos.Normalize())
	fmt.Println(Dot(pos, vel))

	// fmt.Println(pos)
}
