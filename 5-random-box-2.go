package main

import (
	"fmt"
	"math/rand"
	"time"
)

var mapBox = map[string]int{
	"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
}

func main() {
	listOfBox := []string{"A", "B", "C", "D", "E"}
	sumSuccess := 0
	sumFailed := 0

	for i := 0; i < 100; i++ {
		path := []int{}
		loc := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(listOfBox)) + 1
		fmt.Println("LOC", loc)
		path = append(path, loc)
		maxTry := 7
		nyoba := 0
	coba:
		// fmt.Println("pilih box A B C D E?")
		choose := "C"
		nyoba += 1
		if nyoba < maxTry {
			if mapBox[choose] == loc {
				// fmt.Println("CONGRATS")
				sumSuccess += 1
			} else {
				// fmt.Println("hehe", hehe)
				if loc == 1 {
					loc += 1
				} else if loc == 5 {
					loc -= 1
				} else {
					newRand := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
					if newRand == 1 {
						loc += 1
					} else if newRand == 0 {
						loc -= 1
					}
				}
				path = append(path, loc)
				// fmt.Println("LOC", loc)
				goto coba
			}
		} else {
			fmt.Println("GAME OVER", path)
			sumFailed += 1
		}
	}

	fmt.Println("Syccess", sumSuccess)
	fmt.Println("Failed", sumFailed)
}
