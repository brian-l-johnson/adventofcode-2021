package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	distance := 0
	depth := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		command := s[0]
		value, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Errorf(err.Error())
		}
		switch command {
		case "up":
			depth = depth - value
		case "down":
			depth = depth + value
		case "forward":
			distance = distance + value
		}
	}
	fmt.Printf("horizotnal: %v\n", distance)
	fmt.Printf("vertical: %v\n", depth)
	answer := distance * depth
	fmt.Println(answer)
}
