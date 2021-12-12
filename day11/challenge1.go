package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	steps := 100

	scanner := bufio.NewScanner(os.Stdin)
	var og OctopusGrid
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, "")
		var os []octopus
		for _, levelString := range levels {
			l, e := strconv.Atoi(levelString)
			if e != nil {
				panic("invalid level specification")
			}
			os = append(os, newOctopus(l))
		}
		og.grid = append(og.grid, os)
	}

	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			fmt.Printf("%d", og.grid[i][j].energy)
		}
		fmt.Println()
	}

	//fmt.Printf("%v\n", og.getNeighbours(0, 10))

	flashCount := 0

	for i := 0; i < steps; i++ {
		for x := 0; x < len(og.grid); x++ {
			for y := 0; y < len(og.grid[x]); y++ {
				og.grid[x][y].updateEnergy()
			}
		}
		hasflashed := true
		for hasflashed {
			hasflashed = false
			for x := 0; x < len(og.grid); x++ {
				for y := 0; y < len(og.grid[x]); y++ {
					if og.grid[x][y].flashIfAble() {
						hasflashed = true
						flashCount++
						neighbours := og.getNeighbours(x, y)
						//fmt.Printf("%d,%d has the following neighbours: %v\n", x, y, neighbours)
						for _, n := range neighbours {
							og.grid[n.x][n.y].updateEnergy()
						}
					}
				}
			}
		}
		fmt.Printf("on tick %d\n", i)
		for x := 0; x < len(og.grid); x++ {
			for y := 0; y < len(og.grid[x]); y++ {
				fmt.Printf("%d", og.grid[x][y].energy)
			}
			fmt.Println()
		}

		og.resetTick()
	}
	fmt.Printf("%d\n", flashCount)
}

type OctopusGrid struct {
	grid [][]octopus
}

func (og *OctopusGrid) resetTick() {
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			og.grid[i][j].flashed = false
			if og.grid[i][j].energy > 9 {
				og.grid[i][j].energy = 0
			}
		}
	}
}

func (og *OctopusGrid) getNeighbours(x int, y int) []point {
	var ret []point
	if x > 0 {
		ret = append(ret, newPoint(x-1, y))
	}
	if x > 0 && y > 0 {
		ret = append(ret, newPoint(x-1, y-1))
	}
	if x < len(og.grid)-1 {
		ret = append(ret, newPoint(x+1, y))
	}
	if x < len(og.grid)-1 && y < len(og.grid[x])-1 {
		ret = append(ret, newPoint(x+1, y+1))
	}
	if y > 0 {
		ret = append(ret, newPoint(x, y-1))
	}
	if y > 0 && x < len(og.grid)-1 {
		ret = append(ret, newPoint(x+1, y-1))
	}
	if y < len(og.grid[x])-1 {
		ret = append(ret, newPoint(x, y+1))
	}
	if x > 0 && y < len(og.grid[x])-1 {
		ret = append(ret, newPoint(x-1, y+1))
	}
	return ret
}

type octopus struct {
	energy  int
	flashed bool
}

func newOctopus(e int) octopus {
	var o octopus
	o.energy = e
	o.flashed = false
	return o
}

func (o *octopus) updateEnergy() {
	if o.energy <= 9 {
		o.energy++
	}
}

func (o *octopus) flashIfAble() bool {
	if o.energy > 9 && !o.flashed {
		o.flashed = true
		return true
	}
	return false
}

type point struct {
	x int
	y int
}

func newPoint(x int, y int) point {
	var p point
	p.x = x
	p.y = y
	return p
}
