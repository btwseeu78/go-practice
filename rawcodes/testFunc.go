package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusSecond(a, b, c int) int {
	return a + b + c
}

func simple(a, b int) (int, int) {
	return a, b
}

func returnvaridic(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("varidic total: ", total)
}
func Functions() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusSecond(1, 2, 3)
	fmt.Println("1+2+3", res)

	vvalue := 0

	res, vvalue = simple(1, 2)
	fmt.Println(res, vvalue)

	returnvaridic(1, 2, 3, 4, 5, 6)
	returnvaridic(1, 2)

}
