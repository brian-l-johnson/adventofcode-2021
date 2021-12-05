package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func binStringToInt(s string) int64 {
	v, e := strconv.ParseInt(s, 2, 32)
	if e != nil {
		fmt.Println("bad input, can't proceed so dieing")
		os.Exit(1)
	}
	return v
}

func main() {
	var dataBits = 0
	var count = 0

	scanner := bufio.NewScanner(os.Stdin)

	var bitcount []int
	for scanner.Scan() {
		line := scanner.Text()
		if dataBits == 0 {
			dataBits = len(line)
			for i := 0; i < dataBits; i++ {
				bitcount = append(bitcount, 0)
			}
		}
		if len(line) != dataBits {
			fmt.Errorf("data length mismatch!")
			os.Exit(1)
		}
		for i := 0; i < dataBits; i++ {
			if line[i] == '1' {
				bitcount[i]++
			}
		}
		count++
	}

	fmt.Printf("%v", bitcount)

	var g = ""
	var e = ""
	for i := 0; i < dataBits; i++ {
		if bitcount[i] > count/2 {
			g = g + "1"
			e = e + "0"
		} else {
			g = g + "0"
			e = e + "1"
		}
	}

	gi := binStringToInt(g)
	ei := binStringToInt(e)
	fmt.Printf("gamma: %d\n", gi)
	fmt.Printf("epsilon: %d\n", ei)
	fmt.Printf("answer: %d\n", gi*ei)

}
