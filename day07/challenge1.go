package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pos []int
	for scanner.Scan() {
		line := scanner.Text()
		dataStrings := strings.Split(line, ",")
		for _, hpos := range dataStrings {
			v, e := strconv.Atoi(hpos)
			if e != nil {
				panic("invalid horizontal position")
			}
			pos = append(pos, v)
		}
	}
	fmt.Printf("%v\n", pos)
	medianv := median(pos)
	fmt.Printf("median: %d\n", medianv)
	fuelCount := 0
	for _, v := range pos {
		fuelCount = fuelCount + int(math.Abs(float64(medianv)-float64(v)))
	}
	fmt.Println(fuelCount)
}

func median(a []int) int {
	sort.Ints(a)
	fmt.Println(a)
	l := len(a)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		fmt.Println(l / 2)
		fmt.Println(a[(l / 2)])
		fmt.Println(a[(l/2)-1])
		return (a[(l/2)] + a[(l/2)-1]) / 2
	} else {
		return a[l]
	}
}
