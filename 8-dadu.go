package main

import "fmt"

func main() {
	n := 11

	ok := 0
	for x := 1; x <= 6; x++ {
		if x > n-2 {
			break
		}
		for y := 1; y <= 6; y++ {
			if x+y > n-1 {
				break
			}
			for z := 1; z <= 6; z++ {
				if x+y+z > n {
					break
				}
				if x+y+z == n {
					ok += 1
					fmt.Printf("x:: %d, y:: %d, z:: %d \n", x, y, z)
				}
			}
		}
	}
	fmt.Println("ok", ok)
	fmt.Printf("%v/%v", ok, (6 * 6 * 6))
}
