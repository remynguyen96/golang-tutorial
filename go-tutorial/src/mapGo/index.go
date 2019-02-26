package mapGo

func MainMap() {
	/*	myMap := make(map[string]int)
		var myMap2 map[string]int

		fmt.Println(myMap, myMap == nil)
		fmt.Println(myMap2, myMap2 == nil)*/

	/*	firstMap := map[string]int {
			"first": 10,
			"second": 20,
		}
		secondMap := firstMap
		fmt.Println(firstMap, "firstMap")
		secondMap["second222222"] = 50
		fmt.Println(firstMap, "firstMap")
		fmt.Println(secondMap, "secondMap")*/

	/*	findMap := map[string]int {
			"first": 10,
			"second": 20,
		}
		value, found := findMap["fourth"]
		if !found {
			fmt.Println("No found value, please try again!")
		}
		fmt.Println(value, found)*/

	/*	x := 10
		zero(&x)
		fmt.Println(x, "x")*/

	/*	var x *int
		var y int = 10
		x = &y
		fmt.Println(&y == x, "x")
		fmt.Println(*x, "x")
		fmt.Println(y, "y")
		*x = 15
		fmt.Println(y, x)*/

	/*	valTest := new(int)
		one(valTest)
		fmt.Println(*valTest)*/

	/*	a := 10.5
		var pointer *float64
		pointer = &a
		fmt.Println(*pointer, "pointer")
		fmt.Printf("%T", pointer)*/

/*	array := [3]int{1, 2, 3}
	var pointer *[3]int
	pointer = &array
	fmt.Println(pointer, "*pointer")
	fmt.Printf("%T", pointer)*/

/*	value := 30
	var firstPointer *int = &value
	applyPointer(firstPointer)
	fmt.Println(value, "value")
	fmt.Println(*firstPointer, "value")*/

}

func applyPointer(pointer *int) {
	*pointer = 99
}

func zero(x *int) {
	*x = 0
}

func one(firstVal *int) {
	*firstVal = 1
}
