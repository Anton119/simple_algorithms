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
	for i := 0; i < len(B); i++ {
		found := false
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				found = true
				break
			}
		}
		if found {
			C = append(C, A[i])
		}
	}
	A = dif(A, C)
	B = dif(B, C)
	for i := 0; i < len(B); i++ {
		A = append(A, B[i])
	}
	A = sort(A)
	for i := 0; i < len(A); i++ {
		str = append(str, strconv.Itoa(A[i]))
	}
	result = strings.Join(str, ", ")
	fmt.Println(result)
}

func contains(arr []int, elem int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}

func dif(arrA, arrB []int) []int {
	var result []int
	for i := 0; i < len(arrA); i++ {
		if !contains(arrB, arrA[i]) {
			result = append(result, arrA[i])
		}
	}
	return result
}

func sort(arr []int) []int {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				m := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = m
			}
		}
	}
	return arr
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
