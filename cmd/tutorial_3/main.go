package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocatoin: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocatoin: %v", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

// func basic() {
// 	intArr := [3]int32{1, 2, 3}
// 	fmt.Println(intArr)

// 	var intSlice []int32 = []int32{4, 5, 6}
// 	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
// 	intSlice = append(intSlice, 7)
// 	fmt.Printf("\nThe length is %v with capacoty %v\n", len(intSlice), cap(intSlice))

// 	var intSlice2 []int32 = []int32{8, 9}
// 	intSlice = append(intSlice, intSlice2...)
// 	fmt.Println(intSlice)

// 	var intSlice3 []int32 = make([]int32, 3, 8)
// 	fmt.Println(intSlice3)

// 	var myMap map[string]uint8 = make(map[string]uint8)
// 	fmt.Println(myMap)

// 	var myMap2 = map[string]uint8{"John": 22, "Constantine": 40}
// 	fmt.Println(myMap2["Constantine"])

// 	var age, ok = myMap2["John"]
// 	delete(myMap, "Constantine")
// 	if ok {
// 		fmt.Printf("The age is %v", age)
// 	} else {
// 		fmt.Println("Invalid Name")
// 	}

// 	for name, age := range myMap2 {
// 		fmt.Printf("Name: %v, Age: %v \n", name, age)
// 	}
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }
