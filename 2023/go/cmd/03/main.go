package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/simoncrowe/aoc/2023/go/internal/collections"
	"github.com/simoncrowe/aoc/2023/go/internal/grids"
	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

func main() {
	lines := []string{}
	file, err := os.Open("../input/03-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	hoods := grids.BuildNeighbourhoods(lines)

	numbers := grids.ExtractNumbers(hoods)
	partNums := parseSymbolAdjacentNums(numbers)
	fmt.Println("Part 1 answer: ", mathutil.Sum(partNums))

	gearNums := parseGearAdjacentNums(hoods)
	gearTotal := 0
	for _, nums := range gearNums {
		gearTotal += nums[0] * nums[1]
	}
	fmt.Println("Part 2 answer: ", gearTotal)
}

func getDigits() collections.Set[string] {
	digits := make(collections.Set[string])
	digits.Add("0")
	digits.Add("1")
	digits.Add("2")
	digits.Add("3")
	digits.Add("4")
	digits.Add("5")
	digits.Add("6")
	digits.Add("7")
	digits.Add("8")
	digits.Add("9")
	return digits
}

func containsSymbol(hood grids.Neighbourhood) bool {
	symbols := make(collections.Set[string])
	symbols.Add("#")
	symbols.Add("$")
	symbols.Add("%")
	symbols.Add("&")
	symbols.Add("*")
	symbols.Add("+")
	symbols.Add("-")
	symbols.Add("/")
	symbols.Add("=")
	symbols.Add("@")

	for _, char := range hood.Neighbours {
		if symbols.Contains(string(char)) {
			return true
		}
	}
	return false
}

func parseSymbolAdjacentNums(numbers [][]grids.Neighbourhood) []int {
	numText := []string{}
	for _, num := range numbers {
		symbolAdjacent := false
		chars := []rune{}
		for _, hood := range num {
			if containsSymbol(hood) {
				symbolAdjacent = true
			}
			chars = append(chars, rune(hood.Center[0]))
		}
		if symbolAdjacent {
			numText = append(numText, string(chars))
		}
	}
	return parseutil.ParseInts(numText)
}

func parseGearAdjacentNums(hoods [][]grids.Neighbourhood) [][]int {
	digits := getDigits()
	numPairs := [][]int{}

	for y, row := range hoods {
		for x, hood := range row {
			if hood.Center != "*" {
				continue
			}

			numCount := 0
			for _, char := range hood.Neighbours {
				if digits.Contains(string(char)) {
					numCount++
				}
			}
			if numCount >= 2 {
				numbers := grids.GetAdjacentNumbers(x, y, hoods)
				if len(numbers) == 2 {
					pair := grids.ParseNums(numbers)
					numPairs = append(numPairs, pair)
				}
			}
		}
	}
	return numPairs
}
