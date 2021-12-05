package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxX := 0
	maxY := 0

	var pointList []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		p1, p2 := parsePoints(line)
		if p1.x > maxX {
			maxX = p1.x
		}
		if p2.x > maxX {
			maxX = p2.x
		}
		if p1.y > maxY {
			maxY = p1.y
		}
		if p2.y > maxY {
			maxY = p2.y
		}

		if isStraightLine(p1, p2) {
			points := connectPoints(p1, p2)
			pointList = append(pointList, points...)
		}
	}
	fmt.Printf("max x value %d\n", maxX)
	fmt.Printf("max y value %d\n", maxY)
	maxX++
	maxY++

	dangerZone := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		dangerZone[i] = make([]int, maxY)
	}

	for i := 0; i < len(pointList); i++ {
		dangerZone[pointList[i].x][pointList[i].y]++
	}

	count := 0
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			fmt.Printf("%v ", dangerZone[i][j])
			if dangerZone[i][j] > 1 {
				count++
			}
		}
		fmt.Println()
	}
	fmt.Println(count)

}

type Point struct {
	x int
	y int
}

func new(x int, y int) Point {
	var p Point
	p.x = x
	p.y = y
	return p
}

func isStraightLine(p1 Point, p2 Point) bool {
	if p1.x == p2.x || p1.y == p2.y {
		return true
	}
	return false
}

func parsePoints(s string) (Point, Point) {
	pointsString := strings.Split(s, "->")
	p1 := parsePoint(strings.TrimSpace(pointsString[0]))
	p2 := parsePoint(strings.TrimSpace(pointsString[1]))
	return p1, p2
}

func parsePoint(s string) Point {
	ps := strings.Split(s, ",")
	x, e := strconv.Atoi(ps[0])
	if e != nil {
		panic("invalid point spec")
	}
	y, e := strconv.Atoi(ps[1])
	if e != nil {
		fmt.Printf("here: --%v--\n", ps[1])
		panic("invalid point spec")
	}
	var p Point
	p.x = x
	p.y = y
	return p
}

func connectPoints(p1 Point, p2 Point) []Point {
	var points []Point
	fmt.Printf("%d\n", p1)
	fmt.Printf("%d\n", p2)
	if p1.x == p2.x {
		if p1.y > p2.y {
			temp := p2
			p2 = p1
			p1 = temp
		}
		for i := p1.y; i <= p2.y; i++ {
			var p Point
			p.x = p1.x
			p.y = i
			points = append(points, p)
		}
	} else if p1.y == p2.y {
		if p1.x > p2.x {
			temp := p2
			p2 = p1
			p1 = temp
		}
		for i := p1.x; i <= p2.x; i++ {
			var p Point
			p.y = p1.y
			p.x = i
			points = append(points, p)
		}
	} else {
		panic("points aren't in a straint line")
	}
	fmt.Printf("%v\n", points)
	return points
}
