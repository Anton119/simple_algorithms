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
	health, items := ReadInput()
	for _, el := range items {
		if el == "Зелье" {
			health += 10
		}
		if health > 100 {
			health = 100
		}
	}

	fmt.Println(health)
}

func ReadInput() (int, []string) {
	var health int
	var items []string
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	health, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return health, items
}
