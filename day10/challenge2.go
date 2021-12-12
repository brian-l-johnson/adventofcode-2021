package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	openChars := "([{<"
	cloaeChars := ")]}>"

	scores := [4]int{1, 2, 3, 4}
	score := 0
	var scoreList []int

	for scanner.Scan() {
		line := scanner.Text()
		scoreLine := true

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
						scoreLine = false
						break
					}
				}
			}
		}
		if !s.isEmpty() && scoreLine {
			score = 0
			fmt.Printf("%v\n", s)
			for !s.isEmpty() {
				c, _ := s.pop()
				score = (score * 5) + scores[strings.Index(openChars, c)]
			}
			fmt.Printf("%v\n", score)
			scoreList = append(scoreList, score)
		}
	}
	sort.Ints(scoreList)
	sl := len(scoreList)
	fmt.Printf("%v\n", scoreList)
	fmt.Printf("%d\n", scoreList[sl/2])
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
