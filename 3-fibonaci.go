package main

import "fmt"

var sumOdd = 0
var sumEven = 0

func main() {
	x := 2
	y := 3
	n := 3
	i := 0

	sum(x)
	sum(y)

	for i < n-2 {
		temp := x + y
		x = y
		y = temp

		sum(temp)

		i++
	}

	fmt.Println("sumOdd", sumOdd)
	fmt.Println("sumEven", sumEven)
}

func sum(n int) {
	if n%2 == 0 {
		sumEven += n
	} else {
		sumOdd += n
	}
}
