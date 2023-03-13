package main

import "fmt"

func main() {
	strInput := "selamat malam"
	strCharCount := map[string]int{}

	for i := 0; i < len(strInput); i++ {
		str := string(strInput[i])

		fmt.Println(str)
	}

	for _, val := range strInput {
		strCharCount[string(val)]++
	}

	fmt.Println(strCharCount)
}
