package main

import "fmt"

func main() {
	result := SimpleNum(11, 50)

	fmt.Println(result)
}

func SimpleNum(a, b int) []int {
	var count int
	arr := []int{}

	for i := a; i <= b; i++ {

		count = 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				count += 1
			}
		}
		if count == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}
