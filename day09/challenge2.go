package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	scanner := bufio.NewScanner(os.Stdin)

	var fm floorMap

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var ps []floorPoint
		for _, v := range strings.Split(line, "") {
			vi, e := strconv.Atoi(v)
			if e != nil {
				panic("invalid point spec")
			}
			ps = append(ps, newFloorPoint(vi))
		}
		fm.grid = append(fm.grid, ps)
		i++
	}
	fmt.Printf("%v\n", fm.grid)
	var basins []int
	fmt.Printf("grid length: %v\n", len(fm.grid))
	for i := 0; i < len(fm.grid); i++ {
		for j := 0; j < len(fm.grid[i]); j++ {
			if fm.grid[i][j].value != 9 && !fm.grid[i][j].inBasin {
				count := fm.findBasin(i, j, 1)
				fmt.Println(count)
				basins = append(basins, count)
			}
		}
	}

	for i := 0; i < len(fm.grid); i++ {
		for j := 0; j < len(fm.grid[i]); j++ {
			style := colorReset
			if fm.grid[i][j].inBasin {
				style = colorRed
			}
			fmt.Printf("%v%d", style, fm.grid[i][j].value)
		}
		fmt.Println()
	}

	sort.Ints(basins)
	fmt.Printf("%v\n", basins)

	basenLength := len(basins)
	answer := basins[basenLength-1] * basins[basenLength-2] * basins[basenLength-3]
	fmt.Println(answer)

}

type floorMap struct {
	grid [][]floorPoint
}

func (fm *floorMap) findBasin(x int, y int, count int) int {
	bn := count
	if fm.grid[x][y].value != 9 {
		fm.grid[x][y].inBasin = true
		neighbours := fm.getNeighbours(x, y)
		for i := 0; i < len(neighbours); i++ {
			p := neighbours[i]
			if fm.grid[p.x][p.y].value != 9 && fm.grid[p.x][p.y].inBasin == false {
				bn += fm.findBasin(neighbours[i].x, neighbours[i].y, 1)
			}
		}
	}
	return bn
}

func (fm *floorMap) getNeighbours(x int, y int) []point {
	var ret []point
	if x > 0 {
		ret = append(ret, newPoint(x-1, y))
	}
	if x < len(fm.grid)-1 {
		ret = append(ret, newPoint(x+1, y))
	}
	if y > 0 {
		ret = append(ret, newPoint(x, y-1))
	}
	if y < len(fm.grid[x])-1 {
		ret = append(ret, newPoint(x, y+1))
	}
	return ret
}

type floorPoint struct {
	value   int
	inBasin bool
}

func newFloorPoint(v int) floorPoint {
	var fp floorPoint
	fp.value = v
	fp.inBasin = false
	return fp
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
