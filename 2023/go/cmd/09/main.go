package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

func main() {
	seqs := [][]int{}
	file, err := os.Open("../input/09-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seq := parseSeq(scanner.Text())
		seqs = append(seqs, seq)
	}

	totalNext := 0
	for _, seq := range seqs {
		totalNext += predictNext(seq)
	}
	fmt.Println("Part 1 answer: ", totalNext)

	totalPrev := 0
	for _, seq := range seqs {
		totalPrev += predictPrev(seq)
	}
	fmt.Println("Part 2 answer: ", totalPrev)
}

func predictNext(seq []int) int {
	diffs := [][]int{}
	current := seq
	for {
		current = buildDiff(current)
		diffs = append(diffs, current)
		if mathutil.Sum(current) == 0 {
			break
		}
	}
	addend := getAddend(diffs)
	next := seq[len(seq)-1] + addend
	return next
}

func buildDiff(seq []int) []int {
	diff := make([]int, len(seq)-1)
	for i := 1; i < len(seq); i++ {
		diff[i-1] = seq[i] - seq[i-1]
	}
	return diff
}

func getAddend(diffs [][]int) int {
	addend := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		level := diffs[i]
		addend += level[len(level)-1]
	}
	return addend
}

func predictPrev(seq []int) int {
	diffs := [][]int{}
	current := seq
	for {
		current = buildRevDiff(current)
		diffs = append(diffs, current)
		log.Print(current)
		if mathutil.Sum(current) == 0 {
			break
		}
	}
	subtrahend := getSubtrahend(diffs)
	log.Print(subtrahend)
	return seq[0] - subtrahend
}

func buildRevDiff(seq []int) []int {
	diff := make([]int, len(seq)-1)
	for i := 1; i < len(seq); i++ {
		diff[i-1] = seq[i] - seq[i-1]
	}
	return diff
}

func getSubtrahend(diffs [][]int) int {
	sub := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		sub = diffs[i][0] - sub
		log.Print(sub)
	}
	return sub
}

func parseSeq(line string) []int {
	rawNums := strings.Split(line, " ")
	return parseutil.ParseInts(rawNums)
}
