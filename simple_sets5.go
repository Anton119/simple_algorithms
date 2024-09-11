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
	A := ReadInput()
	var result string
	var B []int
	var str []string
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			B = append(B, A[i])
		}
	}
	for i := 0; i < len(B)-1; i++ {
		for j := 0; j < len(B)-1; j++ {
			if B[j] > B[j+1] {
				m := B[j+1]
				B[j+1] = B[j]
				B[j] = m
			}
		}
	}
	for i := 0; i < len(B); i++ {
		str = append(str, strconv.Itoa(B[i]))
	}

	result = strings.Join(str, ", ")

	fmt.Println(result)
}

func ReadInput() []int {
	var A []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &A)
	return A
}
