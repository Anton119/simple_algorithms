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
	data1, data2 := ReadInput()
	var result string
	var data3 []int
	var str []string
	for _, first := range data1 {
		found := false
		for _, second := range data2 {
			if first == second {
				found = true
				break
			}
		}
		if !found {
			data3 = append(data3, first)
		}
	}
	if len(data3) == 0 {
		fmt.Println("Массивы одинаковые")
	} else {
		data3 = sort(data3)
		for _, el := range data3 {
			str = append(str, strconv.Itoa(el))
		}
		result = strings.Join(str, ", ")
		fmt.Println(result)
	}
}

func sort(arr []int) []int {
	m := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				m = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = m
			}
		}
	}
	return arr
}

func ReadInput() ([]int, []int) {
	var data1 []int
	var data2 []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &data1)
	json.Unmarshal([]byte(values[1]), &data2)
	return data1, data2
}
