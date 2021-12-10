package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	bottom1 := ""
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "|")
		values := strings.Fields(data[0])

		var glyphs [10][]string
		var canditates235 []string
		var candidates069 []string

		for _, v := range values {
			switch len(v) {
			case 2:
				//symbol for 1
				glyphs[1] = makeSymbol(v)
			case 3:
				//symbol for 7
				glyphs[7] = makeSymbol(v)
			case 4:
				//symbol for 4
				glyphs[4] = makeSymbol(v)
			case 5:
				//symbol for 2,3, or 5
				canditates235 = append(canditates235, v)
			case 6:
				//symbol for 0, 6, or 9
				candidates069 = append(candidates069, v)
			case 7:
				//symbol for 8
				glyphs[8] = makeSymbol(v)
			}
		}
		//figure out glyphs with 6 segments, this can be 0,6 or 9
		//first look for 9
		for i, v := range candidates069 {
			// 9 contains 4, which the other do not
			if strings.Contains(v, glyphs[4][0]) && strings.Contains(v, glyphs[4][1]) && strings.Contains(v, glyphs[4][2]) && strings.Contains(v, glyphs[4][3]) {
				//fmt.Printf("found 9: %v\n", v)
				candidates069 = removeCandidate(candidates069, i)
				glyphs[9] = makeSymbol(v)
				break
			}
		}
		//next look for 6
		for i, v := range candidates069 {
			// 6 does not have both one characters
			if strings.Contains(v, glyphs[1][0]) && !strings.Contains(v, glyphs[1][1]) {
				//fmt.Printf("found 6 %v\n", v)
				bottom1 = glyphs[1][0]
				candidates069 = removeCandidate(candidates069, i)
				glyphs[6] = makeSymbol(v)
				break
			} else if !strings.Contains(v, glyphs[1][0]) && strings.Contains(v, glyphs[1][1]) {
				//fmt.Printf("found 6: %v\n", v)
				bottom1 = glyphs[1][1]
				candidates069 = removeCandidate(candidates069, i)
				glyphs[6] = makeSymbol(v)
				break
			}
		}
		//by process of elimation, 0 is left
		glyphs[0] = makeSymbol(candidates069[0])
		//fmt.Printf("found 0: %v\n", candidates069[0])

		//next we look at glyphs with 5 segments, so 2, 3 and 5
		//first look for 3
		for i, v := range canditates235 {
			//3 contains 1
			if strings.Contains(v, glyphs[1][0]) && strings.Contains(v, glyphs[1][1]) {
				//fmt.Printf("found 3: %v\n", v)
				canditates235 = removeCandidate(canditates235, i)
				glyphs[3] = makeSymbol(v)
			}
		}
		//next look for 5, 5 does not have the bottom character for 1
		for i, v := range canditates235 {
			if strings.Contains(v, bottom1) {
				//fmt.Printf("found 5: %v\n", v)
				canditates235 = removeCandidate(canditates235, i)
				glyphs[5] = makeSymbol(v)
			}
		}
		//by process of elimination we have 2
		glyphs[2] = makeSymbol(canditates235[0])
		//fmt.Printf("found 2: %v\n", canditates235[0])

		//now the answer
		values = strings.Fields(data[1])
		answer := 0
		for p, v := range values {
			for i, c := range glyphs {
				if reflect.DeepEqual(makeSymbol(v), c) {
					answer += i * int(math.Pow10(3-p))
					break
				}
			}
		}
		fmt.Println(answer)
		sum += answer
	}
	fmt.Println(bottom1)
	fmt.Println(sum)
}

func removeCandidate(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func makeSymbol(s string) []string {
	var ret []string
	for _, c := range s {
		ret = append(ret, string(c))
	}
	sort.Strings(ret)
	return ret
}
