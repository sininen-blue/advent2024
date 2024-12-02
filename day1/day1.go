package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(content)

	inputSlice := strings.Split(input, "\n")

	var leftList []string
	var rightList []string
	for item := range inputSlice {
	}

	fmt.Println(tabsplit)
}
