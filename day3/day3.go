package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputSlice := getInput("input.txt")

	for _, i := range inputSlice {
		fmt.Println(i)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
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
