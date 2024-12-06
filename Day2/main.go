package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file := "Day2\\input.txt"
	content, err := os.Open(file)
	handleErr(err)
	defer content.Close()

	var num1, num2, safeReports int

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 1 {
			safeReports++
			continue
		}
		var isSafe = true
		var isIncreasing, isDecreasing = false, false

		for id, _ := range parts {
			if id == 0 {
				num1, err = strconv.Atoi(parts[id])
				handleErr(err)
				continue
			}

			num2, err = strconv.Atoi(parts[id])
			handleErr(err)

			var diff = num1 - num2

			if diff >= 0 {
				isIncreasing = true
			}
			if diff <= 0 {
				isDecreasing = true
			}

			if absInt(diff) > 3 || (isIncreasing && isDecreasing) {
				isSafe = false
				break
			}
			num1 = num2

		}

		if isSafe {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}
