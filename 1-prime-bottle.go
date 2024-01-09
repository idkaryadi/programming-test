package main

import (
	"fmt"
	"math"
)

var listOfPrime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var isPrime = map[int]bool{
	2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true, 23: true, 29: true,
}

func main() {
	randomInt := 100

	sumOfCombination := []int{}
	fmt.Println("sumOfCombination", sumOfCombination)
	for i := len(listOfPrime) - 1; i > len(listOfPrime)-5; i-- {
		fmt.Println("PRIME::", listOfPrime[i])
		modulo := randomInt % listOfPrime[i]
		sumOfBottle := math.Floor(float64(randomInt) / float64(listOfPrime[i]))
		if modulo < 5 {
			fmt.Println("reset")
		}
		fmt.Println("mod random numb", modulo)
		fmt.Println("floor", sumOfBottle)
	}
}

func getTwoPrime(n int) {
	for k, v := range listOfPrime {
		if v > n {
			sisa := n - listOfPrime[k-1]
			fmt.Println("sisa", sisa)
			if isPrime[sisa] {
				fmt.Println("sisa", sisa)
			}
			if sisa%2 == 0 {
				fmt.Println("genap", sisa)
			}
		}
	}
}
