package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	data := ReadInput()
	var result string
	currentChar := data[0]
	count := 0
	for _, char := range data {
		if currentChar == char {
			count += 1
		} else {
			result += fmt.Sprintf("%s%d", currentChar, count)
			currentChar = char
			count = 1
		}
	}

	result += fmt.Sprintf("%s%d", currentChar, count)
	fmt.Println(result)
}

func ReadInput() []string {
	var data []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
