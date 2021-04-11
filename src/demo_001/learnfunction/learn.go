package learn

import "math"

type Shape interface {
	Area() float64
}

//三角形结构体
type Triangle struct {
	Base   float64
	Height float64
}

//矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

//圆形结构体
type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
