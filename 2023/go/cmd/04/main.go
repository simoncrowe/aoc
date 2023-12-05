package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/collections"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

func main() {
	allWin := [][]int{}
	allMine := [][]int{}
	file, err := os.Open("../input/04-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		rowCount += 1
		win, mine := parseLine(scanner.Text())
		allWin = append(allWin, win)
		allMine = append(allMine, mine)
	}

	score := 0
	for i := 0; i < rowCount; i++ {
		matches := countMatches(allWin[i], allMine[i])
		score += calcScore(matches)
	}
	fmt.Println("Part 1 answer: ", score)

	count := 0
	for i := 0; i < rowCount; i++ {
		count += countCards(allWin, allMine, i)
	}
	fmt.Println("Part 2 answer: ", count)
}

func parseLine(line string) ([]int, []int) {
	data := strings.Split(line, ": ")[1]
	cols := strings.Split(data, "|")

	winRaw := cols[0]
	win := []int{}
	for i := 0; i < len(winRaw); i += 3 {
		win = append(win, parseutil.ParseInt(winRaw[i:i+3]))
	}

	mineRaw := cols[1]
	mine := []int{}
	for j := 0; j < len(mineRaw); j += 3 {
		mine = append(mine, parseutil.ParseInt(mineRaw[j:j+3]))
	}

	return win, mine
}

func countMatches(winNums []int, myNums []int) int {
	win := collections.Set[int]{}
	for _, winNum := range winNums {
		win.Add(winNum)
	}
	count := 0
	for _, mine := range myNums {
		if win.Contains(mine) {
			count++
		}
	}
	return count
}

func calcScore(count int) int {
	var score int
	if count > 0 {
		score = 1 << (count - 1)
	} else {
		score = 0
	}
	return score
}

func countCards(win [][]int, mine [][]int, idx int) int {
	matches := countMatches(win[idx], mine[idx])
	cardCount := 1
	for i := idx + 1; i < idx+1+matches; i++ {
		cardCount += countCards(win, mine, i)
	}
	return cardCount
}
