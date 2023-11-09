package main

import "fmt"

func main() {

	num := 8990

	for num != 0 {
		sum := num % 10

		num = num / 10

		fmt.Println(sum)

	}

}

func addDigits(num int) int {

	sum := 0

	for {

		sum += num % 10

		num = num / 10

		if num == 0 {
			if sum < 10 {
				return sum
			}
			num = sum
			sum = 0
		}
	}

	return 0

}
