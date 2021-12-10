package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "|")
		values := strings.Fields(data[1])
		for _, v := range values {
			if len(v) == 2 || len(v) == 4 || len(v) == 3 || len(v) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}
