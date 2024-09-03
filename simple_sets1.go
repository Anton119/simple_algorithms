package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	A, B := ReadInput()
	var C []int
	var result string
	var str []string
	for i := 0; i < len(A); i++ {
		found := false
		for j := 0; j < len(A); j++ {
			if A[i] == B[j] {
				found = true
				break
			}
		}
		if !found {
			C = append(C, A[i])
		}
	}
	sort.Ints(C)
	for i := 0; i < len(C); i++ {
		str = append(str, strconv.Itoa(C[i]))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)

}

func remove(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func ReadInput() ([]int, []int) {
	var A []int
	var B []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &A)
	json.Unmarshal([]byte(values[1]), &B)
	return A, B
}
