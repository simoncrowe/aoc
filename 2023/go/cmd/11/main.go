package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/simoncrowe/aoc/2023/go/internal/grids"
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
	expanded := expand(grid)
	fmt.Print(expanded)
}

func expand(grid [][]string) [][]string {
	for x := len(grid[0]) - 1; x >= 0; x-- {
		found := false
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "#" {
				found = true
				break
			}
		}
		if found {
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
