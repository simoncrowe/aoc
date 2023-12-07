package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

func main() {
	file, err := os.Open("../input/06-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times := parseNums(scanner.Text())
	bigTime := concatNums(scanner.Text())
	scanner.Scan()
	distances := parseNums(scanner.Text())
	bigDistance := concatNums(scanner.Text())

	counts := []int{}
	for i := 0; i < len(distances); i++ {
		time, distance := times[i], distances[i]
		wins := countWins(time, distance)
		counts = append(counts, wins)
	}

	fmt.Println("Part 1 answer: ", mathutil.Product(counts))

	fmt.Println("Part 2 answer: ", countWins(bigTime, bigDistance))
}

func parseNums(line string) []int {
	intExp := regexp.MustCompile("[0-9]+")
	rawInts := intExp.FindAllString(line, -1)
	return parseutil.ParseInts(rawInts)
}

func concatNums(line string) int {
	intExp := regexp.MustCompile("[0-9]+")
	rawInts := intExp.FindAllString(line, -1)
	joined := strings.Join(rawInts, "")
	return parseutil.ParseInt(joined)
}

func countWins(time int, distance int) int {
	count := 0
	for speed := 0; speed < time; speed++ {
		travelled := (time - speed) * speed
		if travelled > distance {
			count += 1
		}
	}
	return count
}
