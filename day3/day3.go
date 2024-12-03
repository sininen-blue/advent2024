package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	x int
	y int
}

func main() {
	inputSlice := getInput("input.txt")

	commands := []Command{}
	for _, line := range inputSlice {
		// mul(,)
		count := 0
		for _, char := range line {
			inputX := ""
			inputY := ""
			switch count {
			case 0:
				if string(char) == "m" {
					count += 1
				} else {
					count = 0
				}
			case 1:
				if string(char) == "u" {
					count += 1
				} else {
					count = 0
				}
			case 2:
				if string(char) == "l" {
					count += 1
				} else {
					count = 0
				}
			case 3:
				if string(char) == "(" {
					count += 1
				} else {
					count = 0
				}
			case 4:
				if string(char) == "," && inputX != "" {
					inputX = ""
					count += 1
				} else {
					count = 0
				}

				if _, err := strconv.Atoi(string(char)); err == nil {
					inputX = inputX + string(char)
					fmt.Println(inputX)
				} else {
					count = 0
				}
			case 5:
				if string(char) == ")" && inputY != "" {
					inputY = ""
					x, err := strconv.Atoi(inputX)
					y, err := strconv.Atoi(inputY)
					if err != nil {
						fmt.Println(err)
						return
					}
					command := Command{x: x, y: y}
					commands = append(commands, command)

					count = 0
				} else {
					count = 0
				}

				if _, err := strconv.Atoi(string(char)); err == nil {
					inputY = inputY + string(char)
				} else {
					count = 0
				}
			}
		}
	}
	fmt.Println(commands)
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
