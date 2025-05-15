package structsmethodsinterfaces

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// https://google.github.io/styleguide/go/decisions.html#receiver-names
func (r Rectangle) Area() float64 {
	// fmt.Printf("BIR - type of r: %s\n", reflect.TypeOf(r))
	return r.Width * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// fmt.Printf("BIR - type of c: %s\n", reflect.TypeOf(c))
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
