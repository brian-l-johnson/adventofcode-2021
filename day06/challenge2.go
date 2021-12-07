package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fishes = make([]*big.Int, 9)
	for i := 0; i < 9; i++ {
		fishes[i] = big.NewInt(0)
	}
	var newFishes = make([]*big.Int, 9)
	for i := 0; i < 9; i++ {
		newFishes[i] = big.NewInt(0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		seedFish := strings.Split(line, ",")
		for _, f := range seedFish {
			v, e := strconv.Atoi(f)
			if e != nil {
				panic("invalid fish spec")
			}
			fishes[v].Add(fishes[v], big.NewInt(1))
			//fmt.Printf("inisitized fish %d with age %d\n", i, v)
		}
	}
	for day := 0; day < 4096; day++ {
		//newFishes := make([]*big.Int, 9)
		newFishes[8] = fishes[0]
		newFishes[7] = fishes[8]
		newFishes[6].Add(fishes[7], fishes[0])
		newFishes[5] = fishes[6]
		newFishes[4] = fishes[5]
		newFishes[3] = fishes[4]
		newFishes[2] = fishes[3]
		newFishes[1] = fishes[2]
		newFishes[0] = fishes[1]
		fishes = newFishes

		fishSum := big.NewInt(0)
		for i := 0; i < 9; i++ {
			fishSum.Add(fishSum, fishes[i])
		}

		//fmt.Printf("day %d, there are %d fish\n", day, fishSum)
	}
	fishSum := big.NewInt(0)
	for i := 0; i < 9; i++ {
		fishSum.Add(fishSum, fishes[i])
	}

	fmt.Printf("there are %d fish\n", fishSum)
	//fmt.Printf("there are %v fishes\n", len(fishes))

}
