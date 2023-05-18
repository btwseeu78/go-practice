package rawcodes

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
	fmt.Printf("the value is %v and location is %p \n ", nextInt(), test1)
	fmt.Println(nextInt())
	fmt.Printf("the value is %v and location is %p \n", nextInt2(), nextInt2)
	fmt.Printf("the value is %v and location is %p \n ", nextInt2(), nextInt2)
	fmt.Println(nextInt2())
	fmt.Println(test1, test2)

}
