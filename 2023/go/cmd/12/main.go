package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

func main() {
	springs := [][]string{}
	counts := [][]int{}
	file, err := os.Open("../input/12-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s, c := parseSprings(scanner.Text())
		springs = append(springs, s)
		counts = append(counts, c)
	}
	fmt.Println(springs)
	fmt.Println(counts)
}

func parseSprings(line string) ([]string, []int) {
	springExp := regexp.MustCompile("[.#]")
	countExp := regexp.MustCompile("[0-9]+")
	springs := springExp.FindAllString(line, -1)
	countsStr := countExp.FindAllString(line, -1)
	counts := parseutil.ParseInts(countsStr)
	return springs, counts
}
