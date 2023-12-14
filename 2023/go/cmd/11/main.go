package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/EliCDavis/vector/vector2"
	"github.com/simoncrowe/aoc/2023/go/internal/grids"
)

func main() {
	lines := []string{}
	file, err := os.Open("../input/11-test-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	grid := grids.Build(lines)
	expanded := expand(grid)
	locs := locateGalaxies(expanded)
	pairs := buildPairs(locs)
	fmt.Println(len(pairs))
}

func expand(grid [][]string) [][]string {
	for x := len(grid[0]) - 1; x >= 0; x-- {
		empty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "#" {
				empty = false
				break
			}
		}
		if empty {
			for y := 0; y < len(grid); y++ {
				grid[y] = slices.Insert(grid[y], x, grid[y][x])
			}
		}
	}

	for y := len(grid) - 1; y >= 0; y-- {
		if !slices.Contains(grid[y], "#") {
			grid = slices.Insert(grid, y, grid[y])
		}
	}
	return grid
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

func buildPairs(locs []vector2.Vector[int]) [][]vector2.Vector[int] {
	pairs := [][]vector2.Vector[int]{}
	for i := 0; i < len(locs)-1; i++ {
		for j:= i+1; j < len(locs); j++ {
			pair := []vector2.Vector[int]{locs[i], locs[j]}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}
