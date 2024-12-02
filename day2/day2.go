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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var inputSlice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}

	count := 0
	for _, row := range inputSlice {
		rowSlice := strings.Split(row, " ")

		safe := true
		prevLevel := -1
		state := ""
		problems := 0
		for _, levelStr := range rowSlice {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				fmt.Println(err)
				return
			}

			// skip first element
			if prevLevel == -1 {
				prevLevel = level
				continue
			}

			// check whether increasing or decreasing
			difference := level - prevLevel
			prevLevel = level
			curState := state
			if difference < 0 {
				state = "dec"
			} else {
				state = "inc"
			}

			// check unsafe
			absDiff := math.Abs(float64(difference))
			if absDiff > 3 || absDiff <= 0 {
				// checks if levels reach thresholds
				problems += 1
				continue
			}

			if curState != state && curState != "" {
				// checks if it stops increasing or decreasing
				problems += 1

				// keep state the same if problem occurs
				state = curState
				continue
			}

			if problems > 1 {
				safe = false
				break
			}
		}
		if safe == true {
			count += 1
		}
	}
	fmt.Println(count)
}
