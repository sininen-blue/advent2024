package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// plan
// simplest is to just pop out the smallest and do that a bunch
// but doing that with slices
// or maybe get a list of

// im stupid, just sort both lists and get the sums lmao

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

	var leftList []int
	var rightList []int
	for _, row := range inputSlice {
		items := strings.Split(row, "   ")

		leftInt, err := strconv.Atoi(items[0])
		rightInt, err := strconv.Atoi(items[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var sum int
	for i := range inputSlice {
		// difference := math.Abs(float64(leftList[i] - rightList[i]))
		// sum += int(difference)

		sum += leftList[i] * count(leftList[i], rightList)
	}

	fmt.Println(sum)
}

func count(x int, y []int) int {
	count := 0

	for _, item := range y {
		if x == item {
			count += 1
		}
	}

	return count
}
