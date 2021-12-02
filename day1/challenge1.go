package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	last := 0
	count := 0

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Errorf(err.Error())
		}
		if i > 0 {
			if val > last {
				count++
			}
		}
		last = val
		i++
	}

	fmt.Println(count)
}
