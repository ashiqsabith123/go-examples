package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := []int{
		7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6,
	}

	fmt.Println(plusOne(arr))

}
func plusOne(digits []int) []int {
	var num string

	for _, v := range digits {
		num += strconv.Itoa(v)

	}
	fmt.Println(num)

	n, _ := strconv.Atoi(num)
	fmt.Println(n)

	n = n + 1

	fmt.Println(n)

	str := strconv.Itoa(n)

	digits = digits[:0]

	for _, j := range str {
		t := string(j)
		h, _ := strconv.Atoi(t)

		digits = append(digits, h)

	}

	return digits
}
