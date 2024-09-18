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
	power, items := ReadInput()
	for _, el := range items {
		if el == "Энергетик" {
			power += 5
		}
		if el == "Кофе" {
			power += 10
		}
		if power > 100 {
			power = 100
		}
	}

	fmt.Println(power)
}

func ReadInput() (int, []string) {
	var power int
	var items []string
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	power, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return power, items
}
