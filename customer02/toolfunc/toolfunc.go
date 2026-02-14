package toolfunc

import (
	"bufio"
	"fmt"
	"strconv"
)

func PromptString(scanner *bufio.Scanner, prompt string,
	defaultVal string) string {
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		return input
	}
	return defaultVal
}

func PromptIntWithDefault(scanner *bufio.Scanner, prompt string,
	defaultVal int) int {
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		if val, err := strconv.Atoi(input); err == nil {
			return val
		}
	}
	return defaultVal
}
