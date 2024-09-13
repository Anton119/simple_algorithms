package main

import (
	"encoding/json"
	"fmt"

	//"sort"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, data := ReadInput()
	var result string
	var res []int
	var str []string
	for _, el := range data {
		res = append(res, n*el)
	}

	res = sort(res)
	for _, st := range res {
		str = append(str, strconv.Itoa(st))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)
}

func sort(arr []int) []int {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func ReadInput() (int, []int) {
	var n int
	var data []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	n, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &data)
	return n, data
}
