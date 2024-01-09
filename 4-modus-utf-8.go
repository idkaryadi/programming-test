package main

import "fmt"

func main() {
	randomString := "loremipsumdolor"

	strList := []string{}
	strObj := make(map[string]int)

	for _, v := range randomString {
		fmt.Println("v", string(v))
		fmt.Println("strObj[string(v)]", strObj[string(v)])
		if strObj[string(v)] == 0 {
			strList = append(strList, string(v))
		}
		strObj[string(v)] += 1
	}
	// modus
	fmt.Println("strObj", strObj)

	// sorting
	fmt.Println("strList", strList)
}
