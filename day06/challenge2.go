package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fishes [9]uint64

	for scanner.Scan() {
		line := scanner.Text()
		seedFish := strings.Split(line, ",")
		for _, f := range seedFish {
			v, e := strconv.Atoi(f)
			if e != nil {
				panic("invalid fish spec")
			}
			fishes[v]++
			//fmt.Printf("inisitized fish %d with age %d\n", i, v)
		}
	}
	for day := 0; day < 256; day++ {
		var newFishes [9]uint64
		newFishes[8] = fishes[0]
		newFishes[7] = fishes[8]
		newFishes[6] = fishes[7] + fishes[0]
		newFishes[5] = fishes[6]
		newFishes[4] = fishes[5]
		newFishes[3] = fishes[4]
		newFishes[2] = fishes[3]
		newFishes[1] = fishes[2]
		newFishes[0] = fishes[1]
		fishes = newFishes

		var fishSum uint64
		for i := 0; i < 9; i++ {
			fishSum += fishes[i]
		}

		fmt.Printf("day %d, there are %d fish\n", day, fishSum)
	}
	var fishSum uint64
	for i := 0; i < 9; i++ {
		fishSum += fishes[i]
	}

	fmt.Printf("there are %d fish\n", fishSum)
	//fmt.Printf("there are %v fishes\n", len(fishes))

}
