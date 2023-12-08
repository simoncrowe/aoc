package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

const (
	Right = "R"
	Left  = "L"
)

func main() {
	nameExp := regexp.MustCompile("[0-9A-Z]{3}")
	edges := map[string]map[string]string{}
	file, err := os.Open("../input/08-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		names := nameExp.FindAllString(scanner.Text(), 3)
		node, left, right := names[0], names[1], names[2]
		edges[node] = map[string]string{Left: left, Right: right}
	}

	singleSteps := solve("AAA", isAllZ, edges, sequence)
	fmt.Println("Part 1 answer: ", singleSteps)

	//parallelSteps := solveParallel(edges, sequence)

	nodeSteps := []int{}
	for node, _ := range edges {
		if strings.HasSuffix(node, "A") {
			leastSteps := solve(node, endsWithZ, edges, sequence)
			nodeSteps = append(nodeSteps, leastSteps)
		}
	}
	lcm := nodeSteps[0]
	for _, steps := range nodeSteps[1:] {
		lcm = mathutil.LCM(lcm, steps)
	}

	fmt.Println("Part 2 answer: ", lcm)

}

func solve(startNode string, pred func(string) bool, edges map[string]map[string]string, sequence string) int {
	curNode := startNode
	var direction string
	var idx int
	var iter int
	for iter = 0; !pred(curNode); iter++ {
		idx = iter % len(sequence)
		direction = string(sequence[idx])
		curNode = edges[curNode][direction]
	}
	return iter
}

func isAllZ(val string) bool {
	return val == "ZZZ"
}

func endsWithZ(val string) bool {
	return strings.HasSuffix(val, "Z")
}

func solveParallel(edges map[string]map[string]string, sequence string) int {
	curNodes := []string{}
	for nodeName, _ := range edges {
		if strings.HasSuffix(nodeName, "A") {
			curNodes = append(curNodes, nodeName)
		}
	}
	var direction string
	var idx int
	var iter int
	var solved bool = false
	for iter = 0; !solved; iter++ {
		idx = iter % len(sequence)
		direction = string(sequence[idx])

		solved = true
		zCount := 0
		for i := 0; i < len(curNodes); i++ {
			node := curNodes[i]
			nextNode := edges[node][direction]
			curNodes[i] = nextNode
			if !strings.HasSuffix(nextNode, "Z") {
				solved = false
			} else {
				zCount++
			}
		}
		//log.Print("Direction: ", direction)
		if zCount > 0 {
			log.Print("Iteration: ", iter/len(sequence), " - Nodes: ", curNodes)
		}
		//log.Print("Z count: ", zCount)
	}
	return iter
}
