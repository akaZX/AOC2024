package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	file := "Day1\\input.txt"
	content, err := os.Open(file)
	handleErr(err)
	defer content.Close()

	var firstNumbers []int
	var secondNumbers []int

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 {
			firstNum, err := strconv.Atoi(parts[0])
			handleErr(err)
			secondNum, err := strconv.Atoi(parts[1])
			handleErr(err)
			firstNumbers = append(firstNumbers, firstNum)
			secondNumbers = append(secondNumbers, secondNum)
		}
	}

	sort.Ints(firstNumbers)
	sort.Ints(secondNumbers)

	var distance int

	for index, _ := range firstNumbers {
		distance += intAbs(firstNumbers[index] - secondNumbers[index])
	}

	fmt.Printf("Total: %d", distance)
}
