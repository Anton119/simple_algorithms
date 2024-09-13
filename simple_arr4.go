package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := ReadInput()
	var result bool
	for _, boolka := range data {
		if boolka {
			result = false
			break
		} else {
			result = true
		}

	}

	fmt.Println(result)
}

func ReadInput() []bool {
	var data []bool
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
