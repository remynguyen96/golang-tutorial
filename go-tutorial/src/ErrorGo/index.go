package ErrorGo

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	when, what := e.When, e.What
	return fmt.Sprintf("At %v, %s", when, what)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func MainError() {
/*	if err := run(); err != nil {
		fmt.Println(err)
	}*/

	//fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

//func Sqrt(x ErrNegativeSqrt) (float64, error) {
//	if ()
//	return 0, nil
//}
