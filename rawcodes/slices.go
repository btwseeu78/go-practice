package main

import "fmt"

func main() {
	fmt.Println("In Main Func")
	arrayOperation()
	sliceOperation()
	learnCap()
	Maps()
	Functions()
	Closure()
}

func arrayOperation() {
	var a [4]int
	a[0] = 1
	fmt.Printf("address for a is : %p", &a)
	noChnageinArray(a)
	fmt.Println(a)
	ptoa := &a
	pointerchnage(ptoa)
	fmt.Println(a)

}

func noChnageinArray(ar [4]int) {
	ar[0] = 2
	fmt.Printf("adress for ar is : %p", &ar)
	temp := &ar
	temp[0] = 5
	fmt.Println(ar)

}

func pointerchnage(ar *[4]int) {
	ar[0] = 4
}

func sliceOperation() {
	var s []int = []int{1, 2, 3, 4}
	a := s[1:3]
	fmt.Printf("a: length %d, capacity: %d, data: %v\n", len(a), cap(a), a)
}

//undertanding capacity of an slice

func learnCap() {
	originalSlice := make([]int, 3, 5)
	originalSlice = append(originalSlice, 5)
	//originalSlice[0] = 6
	fmt.Printf("originalSlice: length %d, capacity: %d, data: %v\n, and location of the slice: %p", len(originalSlice), cap(originalSlice), originalSlice, originalSlice)
	originalSlice = append(originalSlice, 7, 8)
	fmt.Printf("originalSlice: length %d, capacity: %d, data: %v\n, and location of the slice: %p", len(originalSlice), cap(originalSlice), originalSlice, originalSlice)

}
