package main

import "strconv"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numbers {
		if n%2 != 0 {
			println(strconv.Itoa(n) + " is odd")
		}
		if n%2 == 0 {
			println(strconv.Itoa(n) + " is even")
		}
	}
}
