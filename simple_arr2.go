package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := ReadInput()
	result := data[0][0]
	for _, outer := range data {
		for _, inner := range outer {
			if result < inner {
				result = inner
			}
		}
	}

	fmt.Println(result)
}

func ReadInput() [][]int {
	var data [][]int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
