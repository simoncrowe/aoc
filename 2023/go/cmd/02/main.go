package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

const RedLimit = 12
const BlueLimit = 14
const GreenLimit = 13

func main() {
	idExp := regexp.MustCompile("Game (?P<id>[0-9]+)")
	redExp := regexp.MustCompile("(?P<count>[0-9]+) red")
	greenExp := regexp.MustCompile("(?P<count>[0-9]+) green")
	blueExp := regexp.MustCompile("(?P<count>[0-9]+) blue")

	possibleIds := []int{}
	powers := []int{}

	file, err := os.Open("../input/02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		colonSplit := strings.Split(line, ": ")
		name, data := colonSplit[0], colonSplit[1]

		idText := idExp.FindStringSubmatch(name)[1]
		id := parseutil.ParseInt(idText)

		redMax := parseIntMax(data, redExp)
		blueMax := parseIntMax(data, blueExp)
		greenMax := parseIntMax(data, greenExp)

		if redMax <= RedLimit && blueMax <= BlueLimit && greenMax <= GreenLimit {
			possibleIds = append(possibleIds, id)
		}
		powers = append(powers, redMax*blueMax*greenMax)
	}

	fmt.Println("Part 1 answer: ", mathutil.Sum(possibleIds))
	fmt.Println("Part 2 answer: ", mathutil.Sum(powers))
}

func parseIntMax(text string, numExp *regexp.Regexp) int {
	textNums := numExp.FindAllStringSubmatch(text, -1)
	nums := []int{}
	for _, matches := range textNums {
		parsed := parseutil.ParseInts(matches[1:])
		nums = append(nums, parsed...)
	}
	return slices.Max(nums)
}
