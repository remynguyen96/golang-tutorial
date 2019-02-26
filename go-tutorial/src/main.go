package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*var wg sync.WaitGroup

func init() {
	fmt.Println("Start match!")
	rand.Seed(time.Now().UnixNano())
}

func player(name string, count chan int) {
	defer wg.Done()
	for {
		ball, ok := <- count
		if !ok {
			fmt.Printf("%s thắng!\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s đánh hỏng ở lượt đánh thứ %d!\n", name, ball)
			close(count)
			return
		}
		fmt.Printf("Lượt đánh bóng thành công thứ %d: %s\n", ball, name)
		ball++
		count <- ball
	}
}*/

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"adr"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
}

type Locations struct {
	Locs []Location `json:"loc"`
}

func main() {
	// var f interface{}
	var locs Locations
	jsonStr := `{
	"loc":[
			{"name":"ATM Techcombank", "adr":"Chung cư K26, Dương Quảng Hàm, P.7, Gò Vấp","lat":10.832052230834961,"lon":106.68563842773438,"type":"Tiện ích"},
			{"name":"CGV PVT","adr":"Tầng 5 Vincom Plaza 12 Phan Văn Trị, P7, Gò Vấp","lat":10.827040672302246,"lon":106.68864440917969,"type":"Giải trí"}
		]
}`
	json.Unmarshal([]byte(jsonStr), &locs)
	fmt.Println(locs)
	
	/*	count := make(chan int)
		wg.Add(2)
		go player("Djokovic", count)
		go player("Federer", count)
		count <- 1
		wg.Wait()
		fmt.Println("Ván đấu kết thúc!")*/

	// fmt.Println("This is main!")
	// array.MainArray()
	// mapGo.MainMap()
	// StructGo.Start()
	// MethodGo.MainMethod()
	// InterfaceGo.MainInterface()

	// var vehicles = []InterfaceGo.HornSounder{
	// 	InterfaceGo.Car{"Beep"},
	// }

	// for _, val := range vehicles {
	// 	fmt.Println(val, "val")
	// }

	// ErrorGo.MainError()
	// ConcurrencyGo.MainConcurrency()
	// Crawler.MainCrawler()
	// FuncGo.MainFunc()

	/*	runtime.GOMAXPROCS(8)
		ch := make(chan string)
		doneCh := make(chan bool)
		go abcGen(ch)
		go printer(ch, doneCh)

		println("Golang systems:", runtime.GOOS)
		//time.Sleep(100 * time.Millisecond)
		<-doneCh*/

	/*	runtime.GOMAXPROCS(8)
		var waitGrp sync.WaitGroup
		waitGrp.Add(2)
		go func() {
			defer waitGrp.Done()
			time.Sleep(2 * time.Second)
			fmt.Println("Hello Second")
		}()
		go func() {
			defer waitGrp.Done()
			time.Sleep(1 * time.Second)
			fmt.Println("Hello First")
		}()
		fmt.Println("Wait...")
		waitGrp.Wait()
		fmt.Println("Done!")*/

	/*	f()
		fmt.Println("Returned normally from f.")*/

	/* ------------------------------------------- */
	// WebServer()
}

func WebServer() {
	// mux := http.NewServeMux()
	// hdl := http.HandlerFunc(liteHandler)
	// mux.Handle("/", hdl)
	// http.ListenAndServe(":3000", mux)

	http.HandleFunc("/", LocationHandler)
	http.ListenAndServe(":3000", nil)

	/*	http.HandleFunc("/", liteHandler)
		server := &http.Server{
			Addr:           ":3000",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		server.ListenAndServe()*/

}

func LocationHandler(w http.ResponseWriter, r *http.Request) {

}

func liteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chào mừng đến lập trìjnh Go cho web!")

}

/* ------------------------------------------- */
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g first.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/* ------------------------------------------- */

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}

	doneCh <- true
}
