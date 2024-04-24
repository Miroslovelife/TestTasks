package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)
	result := comtuterSclon(num)
	fmt.Println(result)
}

func comtuterSclon(num int) string {
	computerString := " компьютер"
	numToString := strconv.Itoa(num)
	switch num % 10 {
	case 2, 3, 4:
		computerString += "а"
	case 5, 6, 7, 8, 9, 0:
		computerString += "ов"
	default:
		return numToString + computerString
	}

	return numToString + computerString
}
