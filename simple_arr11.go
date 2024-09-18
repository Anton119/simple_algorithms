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
	data1, data2 := ReadInput()
	var result string
	var arr []int
	var str []string
	for i := 0; i < len(data1); i++ {
		arr = append(arr, data1[i]+data2[i])
	}
	for _, el := range arr {
		str = append(str, strconv.Itoa(el))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)
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
