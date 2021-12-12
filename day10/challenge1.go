package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	openChars := "([{<"
	cloaeChars := ")]}>"

	scores := [4]int{3, 57, 1197, 25137}
	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		instructions := strings.Split(line, "")
		var s stack

		for i := 0; i < len(instructions); i++ {
			if strings.Contains(openChars, instructions[i]) {
				s.push(instructions[i])
			} else if strings.Contains(cloaeChars, instructions[i]) {
				c, e := s.pop()
				if e {
					panic("tried to pop from an empty stack")
				} else {
					if strings.Index(openChars, c) != strings.Index(cloaeChars, instructions[i]) {
						fmt.Printf("found syntax error %v doesn't close out %v\n", instructions[i], c)
						score += scores[strings.Index(cloaeChars, instructions[i])]
						break
					}
				}
			}
		}
	}
	fmt.Println(score)
}

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(val string) {
	*s = append(*s, val)
}
func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", true
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, false
}
