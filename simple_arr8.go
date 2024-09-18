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
	n, items := ReadInput()
	var result string
	var arr []string
	for i := 0; i < n; i++ {
		items = append(items, items[i])
	}

	for _, el := range items {
		arr = append(arr, strconv.Itoa(el))
	}
	result = strings.Join(arr, ", ")
	fmt.Println(result)
}

func ReadInput() (int, []int) {
	var n int
	var items []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	n, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return n, items
}
