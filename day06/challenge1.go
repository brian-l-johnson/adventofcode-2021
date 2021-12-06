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
	var fishes []fish
	for scanner.Scan() {
		line := scanner.Text()
		seedFish := strings.Split(line, ",")
		for i, f := range seedFish {
			v, e := strconv.Atoi(f)
			if e != nil {
				panic("invalid fish spec")
			}
			fmt.Printf("created fish number %v with age %v\n", i, v)
			fishes = append(fishes, newSeedFish(v))
		}
	}
	for i := 0; i < 80; i++ {
		for j := 0; j < len(fishes); j++ {
			if fishes[j].tick() {
				fishes = append(fishes, newFish())
			}
		}
		//fmt.Printf("day %d: %v\n", i, fishes)
		fmt.Printf("day %d done, there are %d fish\n", i, len(fishes))
	}
	fmt.Printf("there are %v fishes\n", len(fishes))

}

type fish struct {
	spawnTimer  int
	justSpawned bool
}

func (f *fish) tick() bool {
	if f.justSpawned {
		f.justSpawned = false
	} else {
		f.spawnTimer--
		if f.spawnTimer < 0 {
			f.spawnTimer = 6
			return true
		}
	}
	return false
}

func newFish() fish {
	var f fish
	f.spawnTimer = 8
	f.justSpawned = true
	return f
}

func newSeedFish(age int) fish {
	var f fish
	f.spawnTimer = age
	f.justSpawned = false
	return f
}
