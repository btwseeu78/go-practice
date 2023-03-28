package listpractice

func Sum(numbers []int) int {
	var sum int

	for _, val := range numbers {
		sum += val
	}
	return sum

}

func SumAll(numbers ...[]int) []int {
	var result []int
	for _, value := range numbers {
		result = append(result, Sum(value))
	}
	return result
}

func SumAllTails(numbers ...[]int) []int {
	result := make([]int, len(numbers))

	for i, value := range numbers {
		if len(value) == 0 {
			result[i] = 0
		} else {
			temp := Sum(value[1:])
			result[i] = temp

		}
	}
	return result
}
