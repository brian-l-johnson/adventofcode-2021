package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	window := ring.New(3)
	pwv := 0
	count := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Errorf(err.Error())
		}
		window.Value = val
		window = window.Next()
		wv := 0
		if i == 2 {
			window.Do(func(i interface{}) {
				pwv = pwv + i.(int)
			})
		}
		if i > 2 {
			window.Do(func(i interface{}) {
				wv = wv + i.(int)
			})
			if wv > pwv {
				count++
			}
			pwv = wv
		}
		i++
	}
	fmt.Println(count)
}
