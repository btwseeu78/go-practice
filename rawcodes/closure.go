package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closure() {
	nextInt := intSeq()
	nextInt2 := intSeq()
	test1 := &nextInt
	test2 := &nextInt2
	fmt.Printf("the value is %v and location is %p and func ", nextInt(), nextInt)
	fmt.Println(nextInt())
	fmt.Printf("the value is %v and location is %p ", nextInt(), &nextInt)
	fmt.Println(nextInt2())
	fmt.Println(test1, test2)

}
