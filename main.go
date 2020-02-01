package main

import "fmt"

func SumSlice(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func Sum(numbers ...int) int {
	return SumSlice(numbers)
}

func main() {
	fmt.Println(SumSlice([]int{1,2,3,4,5}))
	fmt.Println(Sum(1,2,3,4,5))
}