package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filterO2Values(s []string, position int) []string {
	var a, b []string
	var c float32

	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		if s[i][position] == '1' {
			a = append(a, s[i])
			c++
		} else {
			b = append(b, s[i])
		}
	}
	if c >= float32(len(s))/float32(2) {
		return a
	}
	return b
}

func filterCO2Values(s []string, position int) []string {
	var a, b []string
	var c float32
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if s[i][position] == '1' {
			a = append(a, s[i])
			c++
		} else {
			b = append(b, s[i])
		}
	}
	if c >= float32(len(s))/float32(2) {
		return b
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var o2values []string
	var co2values []string

	var dataBits int = 0

	for scanner.Scan() {
		line := scanner.Text()
		if dataBits == 0 {
			dataBits = len(line)
		} else {
			if len(line) != dataBits {
				fmt.Errorf("input is of a different length!")
				os.Exit(1)
			}
		}
		o2values = append(o2values, line)
		co2values = append(co2values, line)
	}

	for i := 0; i < dataBits; i++ {
		o2values = filterO2Values(o2values, i)
	}
	for i := 0; i < dataBits; i++ {
		co2values = filterCO2Values(co2values, i)
	}

	o2, err := strconv.ParseInt(o2values[0], 2, 32)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("O2: %v\n", o2)
	co2, err := strconv.ParseInt(co2values[0], 2, 32)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("CO2: %v\n", co2)

	answer := co2 * o2
	fmt.Println(answer)

}
