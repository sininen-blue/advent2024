package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputSlice := getInput("input.txt")

	count := 0
	for _, row := range inputSlice {
		rowSlice := strings.Split(row, " ")
		safe, _ := checkSafety(rowSlice)

		if safe {
			count += 1
		}
	}
	fmt.Println(count)
}

func checkSafety(rowSlice []string) (bool, []string) {
	safe := true
	dampenedSlice := []string{}
	row := intSlice(rowSlice)

	prevItem := row[0]
	prevState := ""
	state := ""
	for _, item := range row[1:] {
		diff := item - prevItem
		prevItem = item
		prevState = state
		if diff == 0 {
			state = "flat"
			safe = false
			break
		} else if diff < 0 {
			state = "dec"
		} else if diff > 0 {
			state = "inc"
		}

		if prevState != state && prevState != "" {
			safe = false
			break
		}

		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff <= 0 {
			safe = false
			break
		}
	}

	return safe, dampenedSlice
}

func getInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var output []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func intSlice(slice []string) []int {
	var output []int
	for _, item := range slice {
		x, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println(err)
		}

		output = append(output, x)
	}

	return output
}
