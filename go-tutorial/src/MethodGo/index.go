package MethodGo

import (
	"math"
)

func MainMethod() {
	/*	v := Vertex{3, 4}
		fmt.Println(v.AbsNumber(), "v")*/

	/*	f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Abs())*/

	/*	v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println(v.AbsNumber())*/

	/*	v := Vertex{3, 4}
		v.Scale(2)
		ScaleFunc(&v, 10)*/

	/*	p := &Vertex{4, 3}
		p.Scale(3)
		ScaleFunc(p, 8)
		fmt.Println(p, "p")*/

/*		v := Vertex{3, 4}
		fmt.Println(v.AbsNumber(), "v.AbsNumber()")
		fmt.Println(AbsFunc(v), "AbsFunc(v)")*/

/*	p := &Vertex{4, 3}
	fmt.Println(p.AbsNumber(), "v.AbsNumber()")
	fmt.Println(AbsFunc(*p), "AbsFunc(p)")*/

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) AbsNumber() float64 {
	x, y := v.X, v.Y
	return math.Sqrt(x*x + y*y)
}

func AbsFunc(v Vertex) float64 {
	x, y := v.X, v.Y
	return math.Sqrt(x*x + y*y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
