package array

import "fmt"

func MainArray() {
	//var myArray [3]int
	//myArray[0] = 23

	//arrays := [3]int {1, 2, 3}
	//fmt.Println(arrays)
	//fmt.Println(len(arrays))

	//arrList := [...]string {"class1", "class2", "class3", "class4", "class5"}
	//fmt.Println(arrList)

	/*	for i:=0; i < len(arrList); i++ {
			fmt.Println(arrList[i])
		}*/

	/*	for index.go, value := range arrList {
			fmt.Printf("INDEX=%d VALUE=%s ", index.go, value)
			fmt.Println()
		}*/

	// array 2 chiều [row][col] (hang, cot)
	/*	matrix := [4][2]int {
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
		}
		for i:= 0; i < 4; i++ {
			for j:= 0; j < 2; j++ {
				fmt.Print(matrix[i][j], " ")
			}
			fmt.Println()
		}*/

	/*	array := [5]int {1, 2, 3, 4, 5}
		//slice1 := array[1:3]
		//slice1 := array[:]
		//slice1 := array[2:]
		slice1 := array[:4]
		fmt.Println(slice1)*/

	/*	sliceFirst := []int {1,2,3,4,5,6,7,8,9}
		sliceSecond := sliceFirst
		// sliceSecond := sliceFirst[4:8]
		// sliceSecond := sliceFirst
		sliceSecond[0] = 10
		fmt.Println(sliceSecond)
		fmt.Println(sliceFirst)*/

	/*	arrFirst := [4]int{1,2,3,4}
		slice := arrFirst[:]
		slice[0] = 5
		fmt.Println(slice)
		fmt.Println(arrFirst)*/

	/*	slice := []int {1,2,3,4,5,6,7,8,9}
		sliceFirst := slice[2:6]
		sliceText := slice[:3:7]
		fmt.Println(sliceFirst) // 3, 4, 5, 6
		fmt.Println(cap(sliceFirst)) // 7*/

	/*	slice := make([]int, 2, 5)
		fmt.Println(slice)
		fmt.Println(len(slice))
		fmt.Println(cap(slice))*/

	/*	sliceAppend := []int {1,2,3,4}
		slice := append(sliceAppend, 100)
		fmt.Println(slice)*/

	/*	sliceCopy := []string {"First", "Second", "Third", "Fourth"}
		dest := make([]string, 2)
		number := int(copy(dest, sliceCopy))
		fmt.Println(number)
		fmt.Println(dest)*/

	/*	sliceDelete := []string {"First", "Second", "Third", "Fourth"}
		slice1 := append(sliceDelete[:1], sliceDelete[2:]...)
		fmt.Println(slice1)*/

	/*		y := map[string]int {
				"first": 10,
				"second": 20,
				"third": 30,
			}
			x := make(map[string]int)
			x["firstValue"] = 10
			value, status := x["firstValue"]
			fmt.Println(value, status)
			fmt.Println(y, "y")*/

	/*	element := map[string]map[string]string{
			"Infomation": map[string]string{
				"name": "RemyNguyen",
				"age":  "23",
			},
			"Education": map[string]string{
				"Math":      "10",
				"Physical":  "9",
				"Chemistry": "9.5",
			},
		}
		fmt.Println(element, "element")*/

	/*	defer testDefer()
		fmt.Println(add(1, 2, 3), "add()")*/

	//panicError()

/*	list := []int{1, 2, 3, 4}
	fmt.Println(add(list...))
	change(list...)
	fmt.Println(list, "list")*/

}

func change(list ...int) {
	list[0] = 999
}

func panicError() {
	panic("Error here with panic!")
}

func testDefer() {
	fmt.Println("Defer access successful")
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}

	//args = append(args, 10, 11, 12, 13, 14, 15)
	//fmt.Println(args, "args")
	return total
}


