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
	A, B := ReadInput()
	var result string
	var C []int
	var str []string
	for i := 0; i < len(A); i++ {
		found := false
		for j := 0; j < len(A); j++ {
			if A[i] == B[j] {
				found = true
				break
			}
		}
		if found {
			C = append(C, A[i])
		}

	}
	for i := 0; i < len(C)-1; i++ {
		if C[i] > C[i+1] {
			C[i], C[i+1] = C[i+1], C[i]
		}
	}
	for i := 0; i < len(C); i++ {
		str = append(str, strconv.Itoa(C[i]))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)
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
