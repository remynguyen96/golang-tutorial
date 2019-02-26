package InterfaceGo

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type F float64

type I interface {
	M()
}

type T struct {
	str string
}

type Car struct {
	sound string
}

func (c Car) SoundHorn() {
	fmt.Println(c.sound)
}

type Bike struct {
	sound string
}

func (b Bike) SoundHorn() {
	fmt.Println(b.sound)
}

type HornSounder interface {
	SoundHorn()
}

func MainInterface() {
	/*var vehicles = []HornSounder{
		Bike{sound: "Ring"},
		Car{"Beep"},
	}
	for _, v := range vehicles {
		v.SoundHorn()
	}*/

	/*	var a Abser
		f := MyFloat(-math.Sqrt2)
		v := Vertex{3, 4}
		a = f
		fmt.Println(a.Abs())
		a = &v
		fmt.Println(a.Abs())
	*/

	//var i I = T{"Hello!"}
	//i.M()

	/*	var i I
		i = &T{"First I"}
		describe(i)
		i.M()

		i = F(math.Pi)
		describe(i)
		i.M()*/

	/*	var i I
		var t *T
		i = t
		describe(i)
		i.M()
		i = &T{"hello"}
		describe(i)
		i.M()*/

	/*	var i interface{}
		i = 34
		describeNull(i)
		i = "Hello"
		describeNull(i)*/

	/*	var i interface{} = "Hello"
		s := i.(string)
		s, status := i.(string)
		s, ok := i.(float64)
		fmt.Println(s, status)*/

	/*	do(34)
		do("heel")
		do(false)*/

	/*	a := Person{"Arthur Dent", 42}
		z := Person{"Zaphod Beeblebrox", 9001}
		fmt.Println(z, a)*/

	/*	hosts := map[string]IPAddr {
			"loopback": {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}
		for name, ip := range hosts {
			fmt.Printf("%v: %v\n", name, ip)
		}*/

/*	dog := Dog{}
	var m Movement = dog
	m.move()
	var a Animal = dog
	a.speak()*/

}

type Movement interface {
	move()
}

type Animal interface {
	speak()
}

type Dog struct{}

func (d Dog) move() {
	fmt.Println("Movement Here")
}

func (d Dog) speak() {
	fmt.Println("Speak Here")
}

type IPAddr [4]byte

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describeNull(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (f F) M() {
	fmt.Println(f, "f")
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.str, "t.str")
}

func (v *Vertex) Abs() float64 {
	x, y := v.X, v.Y

	return math.Sqrt(x*x + y*y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
