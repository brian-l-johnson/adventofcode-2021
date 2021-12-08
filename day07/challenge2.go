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
	min := 0
	max := 0
	for scanner.Scan() {
		line := scanner.Text()
		dataStrings := strings.Split(line, ",")
		for _, hpos := range dataStrings {
			v, e := strconv.Atoi(hpos)
			if e != nil {
				panic("invalid horizontal position")
			}
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
			pos = append(pos, v)
		}
	}
	fmt.Printf("%v\n", pos)
	/*
		mincost := math.MaxInt
		mincostpos := 0
		//i'm not proud of this, but it runs in a second.  my solution using the mean ended up being off by one so i probably should have just checked on either side of the mean
		for i := min; i < max; i++ {
			tfc := 0
			for _, p := range pos {
				tfc += calculateFuelCost(i, p)
			}
			if tfc < mincost {
				mincost = tfc
				mincostpos = i
			}
		}
		fmt.Printf("min cost is %d at %d\n", mincost, mincostpos)
	*/
	mean := mean(pos)
	fmt.Printf("mean: %d\n", mean)
	for i := mean - 5; i <= mean+5; i++ {
		fuelCount := 0
		for _, v := range pos {
			fuelCount += calculateFuelCost(i, v)
		}
		fmt.Printf("%d costs %d\n", i, fuelCount)
	}

}

func calculateFuelCost(start int, end int) int {
	distance := int(math.Abs(float64(start - end)))
	fuelCost := (int(math.Pow(float64(distance), 2)) + distance) / 2
	return fuelCost
}

func mean(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(float64(sum) / float64(len(a)))
	return int(math.Round(float64(sum) / float64(len(a))))
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
