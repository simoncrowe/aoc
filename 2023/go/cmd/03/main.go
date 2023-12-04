package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/simoncrowe/aoc/2023/go/internal/collections"
	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

type Neighbourhood struct {
	center     string
	neighbours string
}

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

	hoods := buildNeighbourhoods(lines)

	numbers := extractNumbers(hoods)
	partNums := parseSymbolAdjacentNums(numbers)
	fmt.Println("Part 1 answer: ", mathutil.Sum(partNums))

	gearNums := parseGearAdjacentNums(hoods)
	gearTotal := 0
	for _, nums := range gearNums {
		gearTotal += nums[0] * nums[1]
	}
	fmt.Println("Part 2 answer: ", gearTotal)
}

func buildNeighbourhoods(lines []string) [][]Neighbourhood {
	lineChars := [][]rune{}
	for _, line := range lines {
		lineChars = append(lineChars, []rune(line))
	}
	hoods := [][]Neighbourhood{}

	width := len(lines[0])
	height := len(lines)
	for y := 0; y < height; y++ {
		lineHoods := []Neighbourhood{}
		for x := 0; x < width; x++ {
			hoodChars := []rune{}
			if y < height-1 {
				hoodChars = append(hoodChars, lineChars[y+1][x])
			}
			if x < width-1 && y < height-1 {
				hoodChars = append(hoodChars, lineChars[y+1][x+1])
			}
			if x < width-1 {
				hoodChars = append(hoodChars, lineChars[y][x+1])
			}
			if x < width-1 && y > 0 {
				hoodChars = append(hoodChars, lineChars[y-1][x+1])
			}
			if y > 0 {
				hoodChars = append(hoodChars, lineChars[y-1][x])
			}
			if x > 0 && y > 0 {
				hoodChars = append(hoodChars, lineChars[y-1][x-1])
			}
			if x > 0 {
				hoodChars = append(hoodChars, lineChars[y][x-1])
			}
			if x > 0 && y < height-1 {
				hoodChars = append(hoodChars, lineChars[y+1][x-1])
			}
			center := string(lineChars[y][x])
			neighbours := string(hoodChars)
			hood := Neighbourhood{center, neighbours}
			lineHoods = append(lineHoods, hood)
		}
		hoods = append(hoods, lineHoods)
	}
	return hoods
}

func getDigits() collections.StrSet {
	digits := make(collections.StrSet)
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

func extractNumbers(hoods [][]Neighbourhood) [][]Neighbourhood {
	digits := getDigits()

	numbers := [][]Neighbourhood{}
	var curNum []Neighbourhood
	for _, hoodRow := range hoods {
		for _, hood := range hoodRow {
			if digits.Contains(hood.center) {
				if curNum == nil {
					curNum = []Neighbourhood{}
				}
				curNum = append(curNum, hood)
			} else {
				if curNum != nil {
					numbers = append(numbers, curNum)
					curNum = nil
				}
			}
		}
		if curNum != nil {
			numbers = append(numbers, curNum)
			curNum = nil
		}
	}
	return numbers
}

func containsSymbol(hood Neighbourhood) bool {
	symbols := make(collections.StrSet)
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

	for _, char := range hood.neighbours {
		if symbols.Contains(string(char)) {
			return true
		}
	}
	return false
}

func parseSymbolAdjacentNums(numbers [][]Neighbourhood) []int {
	numText := []string{}
	for _, num := range numbers {
		symbolAdjacent := false
		chars := []rune{}
		for _, hood := range num {
			if containsSymbol(hood) {
				symbolAdjacent = true
			}
			chars = append(chars, rune(hood.center[0]))
		}
		if symbolAdjacent {
			numText = append(numText, string(chars))
		}
	}
	return parseutil.ParseInts(numText)
}

func parseGearAdjacentNums(hoods [][]Neighbourhood) [][]int {
	digits := getDigits()
	numPairs := [][]int{}

	for y, row := range hoods {
		for x, hood := range row {
			if hood.center != "*" {
				continue
			}

			numCount := 0
			for _, char := range hood.neighbours {
				if digits.Contains(string(char)) {
					numCount++
				}
			}
			if numCount >= 2 {
				numbers := getAdjacentNumbers(x, y, hoods)
				if len(numbers) == 2 {
					pair := parseNums(numbers)
					numPairs = append(numPairs, pair)
				}
			}
		}
	}
	return numPairs
}

func getAdjacentNumbers(x int, y int, hoods [][]Neighbourhood) [][]Neighbourhood {
	digits := getDigits()

	width := len(hoods[0])
	height := len(hoods)

	nums := [][]Neighbourhood{}

	if y < height-1 && digits.Contains(hoods[y+1][x].center) {
		nums = append(nums, extractNum(x, y+1, hoods))
	}
	if x < width-1 && y < height-1 && digits.Contains(hoods[y+1][x+1].center) {
		nums = append(nums, extractNum(x+1, y+1, hoods))
	}
	if x < width-1 && digits.Contains(hoods[y][x+1].center) {
		nums = append(nums, extractNum(x+1, y, hoods))
	}
	if x < width-1 && y > 0 && digits.Contains(hoods[y-1][x+1].center) {
		nums = append(nums, extractNum(x+1, y-1, hoods))
	}
	if y > 0 && digits.Contains(hoods[y-1][x].center) {
		nums = append(nums, extractNum(x, y-1, hoods))
	}
	if x > 0 && y > 0 && digits.Contains(hoods[y-1][x-1].center) {
		nums = append(nums, extractNum(x-1, y-1, hoods))
	}
	if x > 0 && digits.Contains(hoods[y][x-1].center) {
		nums = append(nums, extractNum(x-1, y, hoods))
	}
	if x > 0 && y < height-1 && digits.Contains(hoods[y+1][x-1].center) {
		nums = append(nums, extractNum(x-1, y+1, hoods))
	}

	uniqueNums := [][]Neighbourhood{}
	for _, num := range nums {
		unique := true
		for _, uniqueNum := range uniqueNums {
			if reflect.DeepEqual(num, uniqueNum) {
				unique = false
			}
		}
		if unique {
			uniqueNums = append(uniqueNums, num)
		}
	}
	return uniqueNums
}

func extractNum(x int, y int, hoods [][]Neighbourhood) []Neighbourhood {
	digits := getDigits()
	width := len(hoods[0])

	number := []Neighbourhood{}

	xIdx := x
	for digits.Contains(hoods[y][xIdx-1].center) {
		xIdx--
		if xIdx == 0 {
			break
		}
	}

	for digits.Contains(hoods[y][xIdx].center) {
		number = append(number, hoods[y][xIdx])
		xIdx++
		if xIdx >= width {
			break
		}
	}

	return number
}

func parseNums(numbers [][]Neighbourhood) []int {
	numText := []string{}
	for _, num := range numbers {
		chars := []rune{}
		for _, hood := range num {
			chars = append(chars, rune(hood.center[0]))
		}
		numText = append(numText, string(chars))
	}
	return parseutil.ParseInts(numText)
}
