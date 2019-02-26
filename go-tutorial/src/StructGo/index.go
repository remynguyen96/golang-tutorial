package StructGo

import (
	"fmt"
	"math"
)

func Start() {
	/*	var c Circle
		var r Rectangle*/
	/*	c := new(Circle)
		r := new(Rectangle)*/

	/*	c := Circle{x: 0, y: 0, r: 5}
		r := Rectangle{x1: 0, y1: 10, x2: 0, y2: 10}
		fmt.Println(c, "c")
		fmt.Println(r, "r")*/

	/*	c := Circle{0, 0, 5}
		r := Rectangle{0, 0, 10, 10}
		fmt.Println(totalArea(&c, &r))*/

	/*	anonymous := struct {
			id int
			name string
			age int
		} {
			id: 2132,
			name: "Billy",
			age: 23,
		}
		fmt.Println(anonymous, "anonymous")*/

	/*	pointerStruct := &Student {
			id: 1,
			name: "Robin",
		}
		fmt.Println((*pointerStruct).id, pointerStruct.id)*/

	/*	n := NoName{"Good", 23}
		fmt.Println(n, "n")*/

	/*	firstStudent := Info{
			id:   23,
			name: "Remy Nguyen",
			subject: Subject{
				math:      10,
				chemistry: 9,
			},
		}
		fmt.Println(firstStudent, "firstStudent")*/

	/*	type struct1 struct {
			id   int
			name string
		}
		s1 := struct1{id: 1, name: "A"}
		s2 := struct1{id: 1, name: "A"}
		fmt.Println(s1 == s2, "s1 == s2")*/

	/*	type struct2 struct {
			id   int
			name string
			info map[int]int
		}
		s1 := struct2{id: 1, name: "A", info: map[int]int{ 1: 5}}
		s2 := struct2{id: 1, name: "A", info: map[int]int{ 1: 5}}
		fmt.Println(s1 == s2, "s1 == s2")*/

/*		pos, neg := adder(), adder()
		for i:= 0; i < 4; i++ {
			fmt.Println(pos(i), neg(-2 * i))
			fmt.Println("=====================")
		}*/

}

func adder() func(int) int {
	sum := 0
	fmt.Println(sum, "sumFirst")
	return func(x int) int {
		fmt.Println(x, "x")
		fmt.Println(sum, "sum")
		sum += x
		return sum
	}
}

type NoName struct {
	string
	int
}

type Info struct {
	id      int
	name    string
	subject Subject
}

type Student struct {
	id   int
	name string
}

type Subject struct {
	math, chemistry int
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Shape interface {
	area() float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}
