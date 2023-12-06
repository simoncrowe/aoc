package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/parseutil"
)

type Route struct {
	sourceFrom int
	sourceTo   int
	destFrom   int
	destTo     int
}

func main() {
	order := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	file, err := os.Open("../input/05-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := parseSeeds(scanner.Text())

	maps := map[string]map[string][]Route{}
	var curMap []Route
	var curSource string
	var curDest string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if curMap != nil {
				maps[curSource] = map[string][]Route{}
				maps[curSource][curDest] = curMap
			}
			curMap = []Route{}
		} else if strings.HasSuffix(line, ":") {
			curSource, curDest = parseHeader(line)
		} else {
			route := parseRoute(line)
			curMap = append(curMap, route)
		}
	}
	if curMap != nil {
		maps[curSource] = map[string][]Route{}
		maps[curSource][curDest] = curMap
	}

	locs := []int{}
	for _, seed := range seeds {
		loc := getDest(seed, maps, order)
		locs = append(locs, loc)
	}
	fmt.Println("Part 1 answer: ", slices.Min(locs))

	var minLoc int = 9_223_372_036_854_775_807
	for i := 0; i < len(seeds); i += 2 {
		from, offset := seeds[i], seeds[i+1]
		to := from + offset
		log.Printf("Offset: %d; seeds %d to %d", i, from, to)
		log.Printf("Current min: %d", minLoc)
		for seed := from; seed < to; seed++ {
			loc := getDest(seed, maps, order)
			if loc < minLoc {
				minLoc = loc
			}
		}
	}
	fmt.Println("Part 2 answer: ", minLoc)
}

func parseSeeds(line string) []int {
	seedsRaw := strings.Split(line, ": ")[1]
	seedsSplit := strings.Split(seedsRaw, " ")
	return parseutil.ParseInts(seedsSplit)
}

func parseHeader(line string) (string, string) {
	sourceToDest := strings.Split(line, " ")[0]
	words := strings.Split(sourceToDest, "-")
	source, dest := words[0], words[2]
	return source, dest
}

func parseRoute(line string) Route {
	rawNums := strings.Split(line, " ")
	nums := parseutil.ParseInts(rawNums)
	destFrom, sourceFrom, length := nums[0], nums[1], nums[2]
	route := Route{
		sourceFrom: sourceFrom,
		sourceTo:   sourceFrom + length,
		destFrom:   destFrom,
		destTo:     destFrom + length,
	}
	return route
}

func getDest(start int, maps map[string]map[string][]Route, order []string) int {
	cur := start
	for t := 1; t < len(order); t++ {
		f := t - 1
		from := order[f]
		to := order[t]
		cur = route(maps[from][to], cur)
	}
	return cur
}

func route(routes []Route, source int) int {
	for _, route := range routes {
		if source >= route.sourceFrom && source < route.sourceTo {
			offset := source - route.sourceFrom
			dest := route.destFrom + offset
			return dest
		}
	}
	return source
}

func expandSeedRanges(seeds []int) []int {
	allSeeds := []int{}
	for i := 0; i < len(seeds); i += 2 {
		from, offset := seeds[i], seeds[i+1]
		to := from + offset
		for s := from; s < to; s++ {
			allSeeds = append(allSeeds, s)
		}
	}
	return allSeeds
}
