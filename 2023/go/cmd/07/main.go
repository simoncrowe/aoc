package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/collections"
	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	OnePair      = 2
	HighCard     = 1
)

type Hand struct {
	cards string
	bet   int
}

func main() {
	hands := []Hand{}
	file, err := os.Open("../input/07-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands = append(hands, parseHand(scanner.Text()))
	}

	slices.SortFunc(hands,
		func(i, j Hand) int {
			return cmpHands(i, j, rankTypeBasic, rankCardBasic)
		})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bet
	}
	fmt.Println("Part 1 answer: ", winnings)

	slices.SortFunc(hands,
		func(i, j Hand) int {
			return cmpHands(i, j, RankTypeJoker, rankCardJoker)
		})

	winnings = 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bet
	}
	fmt.Println("Part 2 answer: ", winnings)
}

func cmpHands(handI Hand, handJ Hand, rankHand func(string) int, rankCard func(byte) int) int {
	iType, jType := rankHand(handI.cards), rankHand(handJ.cards)
	if iType != jType {
		return iType - jType
	}
	for k := 0; k < len(handI.cards); k++ {
		iRank, jRank := rankCard(handI.cards[k]), rankCard(handJ.cards[k])
		if iRank == jRank {
			continue
		}
		return iRank - jRank
	}
	return 0
}

func rankTypeBasic(cards string) int {
	unique := collections.Set[rune]{}
	for _, card := range cards {
		unique.Add(card)
	}
	distinct := len(unique)
	maxCount := 0
	counts := map[rune]int{}
	for _, card := range cards {
		if _, ok := counts[card]; ok {
			counts[card] += 1
		} else {
			counts[card] = 1
		}
	}
	for _, count := range counts {
		if maxCount < count {
			maxCount = count
		}
	}
	if distinct == 1 {
		return FiveOfAKind
	} else if distinct == 2 && maxCount == 4 {
		return FourOfAKind
	} else if distinct == 2 && maxCount == 3 {
		return FullHouse
	} else if distinct == 3 && maxCount == 3 {
		return ThreeOfAKind
	} else if distinct == 3 && maxCount == 2 {
		return TwoPair
	} else if distinct == 4 {
		return OnePair
	} else {
		return HighCard
	}
}

func RankTypeJoker(cards string) int {
	unique := collections.Set[rune]{}
	jokerCount := 0
	for _, card := range cards {
		if string(card) == "J" {
			jokerCount += 1
		} else {
			unique.Add(card)
		}
	}
	distinct := len(unique)
	if distinct == 0 && jokerCount > 0 {
		distinct = 1
	}
	maxCount := 0
	counts := map[rune]int{}
	for _, card := range cards {
		if string(card) == "J" {
			continue
		}
		if _, ok := counts[card]; ok {
			counts[card] += 1
		} else {
			counts[card] = 1
		}
	}
	for _, count := range counts {
		if maxCount < count {
			maxCount = count
		}
	}
	maxCount += jokerCount
	if distinct == 1 {
		return FiveOfAKind
	} else if distinct == 2 && maxCount == 4 {
		return FourOfAKind
	} else if distinct == 2 && maxCount == 3 {
		return FullHouse
	} else if distinct == 3 && maxCount == 3 {
		return ThreeOfAKind
	} else if distinct == 3 && maxCount == 2 {
		return TwoPair
	} else if distinct == 4 {
		return OnePair
	} else {
		return HighCard
	}
}

func rankCardBasic(char byte) int {
	switch name := string(char); name {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}
	log.Fatalf("Cannot map card name %s", string(char))
	return 0
}

func rankCardJoker(char byte) int {
	switch name := string(char); name {
	case "J":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}
	log.Fatalf("Cannot map card name %s", string(char))
	return 0
}
func parseHand(line string) Hand {
	cols := strings.Split(line, " ")
	cards := cols[0]
	bet := parseutil.ParseInt(cols[1])
	return Hand{cards, bet}
}
