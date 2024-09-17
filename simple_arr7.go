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
	var str []string
	for i := 0; i < n; i++ {
		items = append(items, items[i])
	}
	for i := 0; i < n; i++ {
		items = deleteEl(items, 0)
	}
	for _, el := range items {
		str = append(str, strconv.Itoa(el))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)
}

func deleteEl(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
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
