package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := ReadInput()
	m := data[0][0]

	for _, outer_array := range data {
		for _, inner_array := range outer_array {
			if m > inner_array {
				m = inner_array
			}
		}
	}
	fmt.Println(m)
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
