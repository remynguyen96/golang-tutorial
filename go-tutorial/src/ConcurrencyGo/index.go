package ConcurrencyGo

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func MainConcurrency() {
	/*	go say("World")
		say("Hello")*/

	/*	s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[:len(s)/2], c) // 7, 2, 8 => 17
		go sum(s[len(s)/2:], c) // -9, 4 ,0 => -5
		x, y := <- c, <- c
		fmt.Println(x, y, x + y)*/

	/*	ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		fmt.Println(<-ch, "<-ch1")
		fmt.Println(<-ch, "<-ch2")*/

	/*	c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c {
			fmt.Println(i, "i")
		}
	*/

/*	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}*/

/*	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
*/
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	c <- sum
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
