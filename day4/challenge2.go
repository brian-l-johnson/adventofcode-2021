package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	i := 0
	var draws []int
	var boards []Board
	var boardString string
	boardReadLineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			drawsStrings := strings.Split(line, ",")
			for j := 0; j < len(drawsStrings); j++ {
				v, e := strconv.Atoi(drawsStrings[j])
				if e != nil {
					panic("invalid value in draw")
				}
				draws = append(draws, v)
			}
		} else {
			if line == "" && i > 1 {
				var b Board
				b.picked = [25]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
				var values [25]int
				b.won = false
				bsf := strings.Fields(boardString)
				if len(bsf) != 25 {
					panic("invalid board size")
				}
				for j := 0; j < 25; j++ {
					v, e := strconv.Atoi(bsf[j])
					if e != nil {
						panic("invalid value in board string")
					}
					values[j] = v
				}
				b.values = values
				fmt.Printf("%v\n", values)
				boards = append(boards, b)
				boardString = ""
			} else {
				boardString = boardString + " " + line
				boardReadLineCount++
			}
		}

		i++
	}
	fmt.Printf("draws: %v\n", draws)

	for i := 0; i < len(draws); i++ {
		fmt.Printf("draw %d, looking for %d\n", i, draws[i])
		for j := 0; j < len(boards); j++ {
			score := boards[j].CheckDraw(draws[i])
			if score > 0 {
				fmt.Printf("Winning board found! %d\n", score*draws[i])
				boards[j].won = true
				//os.Exit(0)
			}
		}
	}

}

type Board struct {
	values [25]int
	picked [25]bool
	won    bool
}

func (b *Board) CheckDraw(draw int) int {
	if b.won {
		return 0
	}
	for i := 0; i < 25; i++ {
		if b.values[i] == draw {
			b.picked[i] = true
		}
	}
	b.printBoard()
	for i := 0; i < 5; i++ {
		win := true
		for j := 0; j < 5; j++ {
			if b.picked[i*5+j] == false {
				win = false
				break
			}
		}
		if win {
			return b.boardScore()
		}
	}
	for i := 0; i < 5; i++ {
		win := true
		for j := 0; j < 5; j++ {
			if b.picked[i+j*5] == false {
				win = false
				break
			}
		}
		if win {
			b.won = true
			return b.boardScore()
		}
	}
	return 0
}

func (b Board) boardScore() int {
	score := 0
	for i := 0; i < 25; i++ {
		if !b.picked[i] {
			fmt.Printf("%d+", b.values[i])
			score += b.values[i]
		}
	}
	fmt.Println()
	return score
}

func (b Board) printBoard() {
	fmt.Println("------------")
	for i := 0; i < 25; i++ {
		if i%5 == 0 {
			fmt.Println()
		}
		if b.picked[i] {
			fmt.Printf(" * ")
		} else {
			fmt.Printf("%2d ", b.values[i])
		}
	}
	fmt.Println("\n------------")
}
