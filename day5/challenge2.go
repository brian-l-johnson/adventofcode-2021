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

		points := connectPoints(p1, p2)
		pointList = append(pointList, points...)
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
	xinc := 0
	yinc := 0
	if p1.x < p2.x {
		xinc = 1
	} else if p1.x > p2.x {
		xinc = -1
	}
	if p1.y < p2.y {
		yinc = 1
	} else if p1.y > p2.y {
		yinc = -1
	}
	connecting := true
	tx := p1.x
	ty := p1.y
	points = append(points, p1)
	for connecting {
		tx = tx + xinc
		ty = ty + yinc
		if tx == p2.x && ty == p2.y {
			connecting = false
		}
		var p Point
		p.x = tx
		p.y = ty
		points = append(points, p)
	}

	fmt.Printf("%v\n", points)
	return points
}
