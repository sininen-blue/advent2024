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
		target := []rune{'m', 'u', 'l', '(', ',', ')'}

		count := 0
		inputX := ""
		inputY := ""

		for _, char := range line {
			if count == 4 {
				if _, err := strconv.Atoi(string(char)); err == nil {
					inputX += string(char)
					continue
				} else if inputX == "" {
					count = 0
					continue
				}
			}

			if count == 5 {
				if _, err := strconv.Atoi(string(char)); err == nil {
					inputY += string(char)
					continue
				} else if inputY == "" {
					count = 0
					continue
				}

				if char == ')' {
					x, err := strconv.Atoi(inputX)
					y, err := strconv.Atoi(inputY)
					if err != nil {
						fmt.Println(err)
						return
					}
					command := Command{x: x, y: y}
					fmt.Println(command)
					commands = append(commands, command)
					count = 0
					inputX = ""
					inputY = ""
				} else {
					inputX = ""
					inputY = ""
					count = 0
					continue
				}
			}

			if char == target[count] {
				// fmt.Println(string(char))
				count += 1
			} else {
				inputX = ""
				inputY = ""
				count = 0
			}
		}
	}
	sum := 0
	for _, com := range commands {
		sum += com.x * com.y
	}

	fmt.Println("total muls:", sum)
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
