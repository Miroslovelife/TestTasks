package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int
	fmt.Scan(&num)
	table := multTable(num)
	for i := 0; i < num; i++ {
		fmt.Println(strings.Join(table[i], " "))
	}
}

func multTable(num int) [][]string {
	table := [][]string{}

	for i := 1; i <= num; i++ {
		row := []string{}
		for j := 1; j <= num; j++ {
			stringRow := strconv.Itoa(i * j)
			row = append(row, stringRow)
		}
		table = append(table, row)
	}

	return table

}
