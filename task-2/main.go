package main

import "fmt"

func main() {
	arr := []int{48, 16, 256, 512}
	result := SimpleNum(arr)

	fmt.Println(result)

}

func SimpleNum(nums []int) []int {

	arr := []int{}
	resultArr := []int{}
	for _, elem := range nums {
		for i := 2; i <= Minimum(nums); i++ {
			if elem%i == 0 {
				arr = append(arr, i)
			}
		}
	}

	var count int
	for _, elem := range arr {
		count = 0
		for i := 0; i < len(arr); i++ {
			if elem == arr[i] {
				count += 1
			}
		}
		if count == len(nums) {
			resultArr = append(resultArr, elem)
		}

	}
	result := RemoveDuplicates(resultArr)

	return result
}

func Minimum(num []int) int {
	minimum := num[0]
	for _, elem := range num {
		if elem < minimum {
			minimum = elem
		}
	}
	return minimum
}

func RemoveDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	var result []int

	for _, item := range arr {
		if _, ok := unique[item]; !ok {
			unique[item] = true
			result = append(result, item)
		}
	}

	return result
}
