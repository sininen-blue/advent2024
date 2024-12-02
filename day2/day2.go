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
				safe = false
				break
			}

			if curState != state && curState != "" {
				// checks if it stops increasing or decreasing
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
