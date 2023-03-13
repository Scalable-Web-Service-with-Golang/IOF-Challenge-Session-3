package main

import "fmt"

func main() {
	input := "selamat malam"
	strCharCount := map[string]int{}

	for i := 0; i < len(input); i++ {
		str := string(input[i])

		fmt.Println(str)
		strCharCount[str]++
	}

	fmt.Println(strCharCount)
}
