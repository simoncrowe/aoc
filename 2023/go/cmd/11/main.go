package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/EliCDavis/vector/vector2"
	"github.com/dominikbraun/graph"
	"github.com/simoncrowe/aoc/2023/go/internal/grids"
	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

func main() {
	lines := []string{}
	file, err := os.Open("../input/11-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	grid := grids.Build(lines)
	locs := locateGalaxies(grid)
	expandedLocs := expandLocations(grid, locs, 1)
	pairs := buildPairs(expandedLocs)
	total := 0
	for _, pair := range pairs {
		total += findShortestPath(pair)
	}
	fmt.Println("Part 1 answer: ", total)

	furtherLocs := locateGalaxies(grid)
	furtherExpandedLocs := expandLocations(grid, furtherLocs, 999_999)
	furtherPairs := buildPairs(furtherExpandedLocs)
	furtherTotal := 0
	for _, pair := range furtherPairs {
		furtherTotal += findShortestPath(pair)
	}
	fmt.Println("Part 2 answer: ", furtherTotal)
}

func locateGalaxies(grid [][]string) []vector2.Vector[int] {
	locs := []vector2.Vector[int]{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "#" {
				locs = append(locs, vector2.New[int](x, y))
			}
		}
	}
	return locs
}

func expandLocations(grid [][]string, locs []vector2.Vector[int], offset int) []vector2.Vector[int] {
	xAddend := vector2.New[int](offset, 0)
	for x := len(grid[0]) - 1; x >= 0; x-- {
		empty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "#" {
				empty = false
				break
			}
		}
		if empty {
			for i, loc := range locs {
				if loc.X() > x {
					locs[i] = loc.Add(xAddend)
				}
			}
		}
	}

	yAddend := vector2.New[int](0, offset)
	for y := len(grid) - 1; y >= 0; y-- {
		if !slices.Contains(grid[y], "#") {
			for i, loc := range locs {
				if loc.Y() > y {
					locs[i] = loc.Add(yAddend)
				}
			}
		}
	}
	return locs
}

func hashLoc(loc vector2.Vector[int]) string {
	return fmt.Sprint(loc)
}

func buildGraph(grid [][]string) graph.Graph[string, vector2.Vector[int]] {
	height := len(grid)
	width := len(grid[0])
	g := graph.New(hashLoc)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			loc := vector2.New[int](x, y)
			g.AddVertex(loc)
			if x > 0 {
				g.AddEdge(hashLoc(loc), hashLoc(vector2.New[int](x-1, y)))
			}
			if y > 0 {
				g.AddEdge(hashLoc(loc), hashLoc(vector2.New[int](x, y-1)))
			}
			if x < width-1 {
				g.AddEdge(hashLoc(loc), hashLoc(vector2.New[int](x+1, y)))
			}
			if y < height-1 {
				g.AddEdge(hashLoc(loc), hashLoc(vector2.New[int](x, y+1)))
			}
		}
	}
	return g
}

func buildPairs(locs []vector2.Vector[int]) [][]vector2.Vector[int] {
	pairs := [][]vector2.Vector[int]{}
	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			pair := []vector2.Vector[int]{locs[i], locs[j]}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func findShortestPath(pair []vector2.Vector[int]) int {
	start, end := pair[0], pair[1]
	diff := start.Sub(end)
	return mathutil.Abs(diff.X()) + mathutil.Abs(diff.Y())
}
