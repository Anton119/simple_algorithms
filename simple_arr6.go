package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	k, data := ReadInput()
	var result bool

	result = FindNumber(k, data)

	fmt.Println(result)
}

func FindNumber(k int, data [][]int) bool {
	for _, first := range data {
		for _, second := range first {
			if k == second {
				return true
			}
		}
	}
	return false
}

func ReadInput() (int, [][]int) {
	var k int
	var data [][]int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	k, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &data)
	return k, data
}
