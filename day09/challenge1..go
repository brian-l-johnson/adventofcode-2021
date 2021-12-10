package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		for j, v := range strings.Split(line, "") {
			vi, e := strconv.Atoi(v)
			if e != nil {
				panic("invalid point spec")
			}
			ps = append(ps, newFloorPoint(i, j, vi))
		}
		fm.grid = append(fm.grid, ps)
		i++
	}
	fmt.Printf("%v\n", fm.grid)
	fmt.Printf("grid length: %v\n", len(fm.grid))
	for i := 0; i < len(fm.grid); i++ {
		for j := 0; j < len(fm.grid[i]); j++ {
			fm.resetVisited()
			min := fm.findLocalMin(i, j)
			fmt.Printf("local min for %d,%d found at %d,%d\n", i, j, min.localmin.x, min.location.y)
			fm.grid[i][j].localmin = min.location
			fm.grid[min.location.x][min.location.y].islocalmin = true
		}
	}
	answer := 0
	for i := 0; i < len(fm.grid); i++ {
		for j := 0; j < len(fm.grid[i]); j++ {
			style := colorReset
			if fm.grid[i][j].islocalmin {
				style = colorRed
				answer += fm.grid[i][j].value + 1
			}
			fmt.Printf("%v%d", style, fm.grid[i][j].value)
		}
		fmt.Println()
	}
	fmt.Printf("answer is %d\n", answer)

}

type floorMap struct {
	grid [][]floorPoint
}

func (fm *floorMap) findLocalMin(x int, y int) floorPoint {
	fp := fm.grid[x][y]
	fm.grid[x][y].visited = true
	fmt.Printf("Checking %d,%d\n", x, y)
	if fp.islocalmin {
		return fp
	}
	if fp.localminfound {
		return fm.grid[fp.localmin.x][fp.location.y]
	}
	var candidateLocalMin []point
	for _, p := range fm.getNeighbours(x, y) {
		cfp := fm.grid[p.x][p.y]
		if cfp.value <= fp.value && !cfp.visited {
			candidateLocalMin = append(candidateLocalMin, fm.findLocalMin(p.x, p.y).location)
		}
	}
	if len(candidateLocalMin) == 0 {
		return fp
	}
	minCandidateValue := math.MaxInt
	minCandidateIndex := 0

	for i := 0; i < len(candidateLocalMin); i++ {
		if fm.grid[candidateLocalMin[i].x][candidateLocalMin[i].y].value < minCandidateValue {
			minCandidateIndex = i
			minCandidateValue = fm.grid[candidateLocalMin[i].x][candidateLocalMin[i].y].value
		}
	}
	if minCandidateValue < fp.value {
		return fm.grid[candidateLocalMin[minCandidateIndex].x][candidateLocalMin[minCandidateIndex].y]
	} else {
		return fp
	}

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

func (fm *floorMap) resetVisited() {
	for i := 0; i < len(fm.grid); i++ {
		for j := 0; j < len(fm.grid[i]); j++ {
			fm.grid[i][j].visited = false
		}
	}
}

type floorPoint struct {
	value         int
	location      point
	visited       bool
	localminfound bool
	localmin      point
	islocalmin    bool
}

func newFloorPoint(x int, y int, v int) floorPoint {
	var fp floorPoint
	fp.value = v
	fp.location = newPoint(x, y)
	fp.visited = false
	fp.islocalmin = false
	fp.localminfound = false
	if v == 0 {
		fp.islocalmin = true
		fp.localmin = newPoint(x, y)
		fp.localminfound = true
	}
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
