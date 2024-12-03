package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	x int
	y int
}

func main() {
	inputSlice := getInput("input.txt")

	commands := []Command{}
	for _, line := range inputSlice {
		active := true
		for i := 0; i < len(line); i++ {
			if active {
				if i+7 < len(line) && parseDont(line[i:i+7]) {
					i = i + 7
					active = false
					fmt.Println("don't()")
					continue
				}
				if i+12 < len(line) {
					command, err := parseMul(line[i : i+12])
					if err != nil {
						if err.Error() == "Invalid" {
							continue
						}
						fmt.Println(err)
						continue
					}
					fmt.Println(command)
					commands = append(commands, command)
				}
			} else {
				if i+4 < len(line) && parseDo(line[i:i+4]) {
					i = i + 4
					active = true
					fmt.Println("do()")
					continue
				}
			}
		}
	}

	sum := 0
	for _, com := range commands {
		sum += com.x * com.y
	}

	fmt.Println("total muls:", sum)
}

func parseDont(line string) bool {
	if line == "don't()" {
		return true
	}
	return false
}

func parseDo(line string) bool {
	if line == "do()" {
		return true
	}
	return false
}

func parseMul(line string) (Command, error) {
	var command Command
	var err error

	if line[0:4] == "mul(" {
		args := parseMulArgs(line[4:])

		argSlice := strings.Split(args, ",")
		if len(argSlice) < 2 {
			return command, errors.New(fmt.Sprintln("Second argument invalid", line))
		}

		if string(line[4+len(args)]) != ")" {
			return command, errors.New(fmt.Sprintln("Not closed properly", line))
		}

		x, err := strconv.Atoi(argSlice[0])
		y, err := strconv.Atoi(argSlice[1])
		if err != nil {
			fmt.Println(err)
		}

		command.x = x
		command.y = y
		return command, err
	} else {
		err = errors.New("Invalid")
		return command, err
	}
}

func parseMulArgs(line string) string {
	var output string
	for i := 0; i < len(line); i++ {
		if _, err := strconv.Atoi(string(line[i])); err == nil {
			output += string(line[i])
		} else if string(line[i]) == "," {
			output += string(line[i])
		}
	}

	return output
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
