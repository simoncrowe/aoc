package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type intConv func(string) (int, error)

func main() {
	digitExp, err := regexp.Compile("[0-9]")
	if err != nil {
		log.Fatal(err)
	}
	digitVals := []int{}

	textExp, err := regexp.Compile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	if err != nil {
		log.Fatal(err)
	}
	textVals := []int{}

	file, err := os.Open("../input/01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digitVal := parseLine(line, digitExp, strconv.Atoi)
		digitVals = append(digitVals, digitVal)
		textVal := parseLine(line, textExp, textToInt)
		textVals = append(textVals, textVal)
	}

	digitSum := 0
	for _, val := range digitVals {
		digitSum += val
	}
	fmt.Println("Part 1 answer: ", digitSum)

	textSum := 0
	for _, val := range textVals {
		textSum += val
	}
	fmt.Println("Part 2 answer: ", textSum)
}

func parseLine(line string, numExp *regexp.Regexp, toInt intConv) int {
	digits := numExp.FindAllString(line, -1)
	ints := []int{}
	for _, digit := range digits {
		val, err := toInt(digit)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, val)
	}
	first, last := ints[0], ints[len(ints)-1]
	strNum := fmt.Sprintf("%d%d", first, last)
	num, err := toInt(strNum)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func textToInt(text string) (int, error) {
	parsed, err := strconv.Atoi(text)
	if err == nil {
		return parsed, nil
	}
	var val int
	var matchErr error
	switch word := text; word {
	case "zero":
		val = 0
	case "one":
		val = 1
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	default:
		matchErr = fmt.Errorf("Could not parse '%s' as digit", text)
	}
	return val, matchErr
}
