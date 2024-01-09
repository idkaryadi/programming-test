package main

import "fmt"

func main() {
	matrix1 := [3][4]int{{1, 3, 3, 4}, {1, 3, 3, 4}, {1, 3, 3, 4}}
	matrix2 := [4][4]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}

	// gak usah validasi ukuran karen mxn dikali nxm

	// jumlah baris matrix 1 as final dimention
	m := len(matrix1)

	// jumlah kolom matrix 2
	n := len(matrix2[0])
	result := [][]int{}
	// i => jumlah baris
	for i := 0; i < m; i++ {
		// generate per baris
		temp := make([]int, n)
		// j => kolom matrix kedua
		for j := 0; j < n; j++ {
			sum := 0
			// k => baris matrix pertama
			for k := 0; k < len(matrix1[j]); k++ {
				// 0 ntar diganti i
				sum += matrix1[i][k] * matrix2[k][j]
			}
			temp[j] = sum
		}
		fmt.Println("temp", temp)
		result = append(result, temp)
	}
	fmt.Println("resylt", result)

}
