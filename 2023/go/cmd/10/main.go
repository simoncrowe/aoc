package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

type Ground struct{ sym string }

type Pipe struct {
	inA, inB mathutil.Vec2
	sym      string
}

type Traversable interface {
	isOpen(mathutil.Vec2) bool
	runsAlongX() bool
	runsAlongY() bool
	getSym() string
}

func (g Ground) isOpen(inDir mathutil.Vec2) bool { return false }
func (g Ground) runsAlongX() bool                { return false }
func (g Ground) runsAlongY() bool                { return false }
func (g Ground) getSym() string                  { return g.sym }

func (p Pipe) isOpen(inDir mathutil.Vec2) bool {
	return inDir == p.inA || inDir == p.inB
}

func (p Pipe) runsAlongY() bool {
	return p.inA.Y != 0 || p.inB.Y != 0
}

func (p Pipe) runsAlongX() bool {
	return p.inA.X != 0 || p.inB.X != 0
}

func (p Pipe) getOutDir(inDir mathutil.Vec2) mathutil.Vec2 {
	var out mathutil.Vec2
	if inDir == p.inA {
		out = p.inB.Flip()
	} else if inDir == p.inB {
		out = p.inA.Flip()
	} else {
		log.Fatalln("Pipe ", p, " not open to dir ", inDir)
	}
	return out
}

func (p Pipe) getSym() string   { return p.sym }

func main() {
	lines := []string{}
	file, err := os.Open("../input/10-test-input-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	slices.Reverse(lines)

	grid, start := makeTraversable(lines)
	loopLocs, loopPipes := walkLoop(grid, start)
	log.Print(loopPipes)
	fmt.Println("Part 1 answer :", len(loopPipes)/2)
	
	enclosedLocs, enclosedObjs := getEnclosed(loopLocs, loopPipes, grid)
	log.Print(enclosedLocs)
	log.Print(enclosedObjs)
	fmt.Println("Part 2 answer :", len(enclosedLocs))
}

func walkLoop(grid [][]Traversable, start mathutil.Vec2) ([]mathutil.Vec2, []Pipe) {
	startPipe := grid[start.Y][start.X].(Pipe)
	dir := startPipe.inA.Flip()
	pipes := []Pipe{startPipe}
	locs := []mathutil.Vec2{start}
	loc := start.Add(dir)
	for loc != start {
		locs = append(locs, loc)
		pipe := grid[loc.Y][loc.X].(Pipe)
		pipes = append(pipes, pipe)
		dir = pipe.getOutDir(dir)
		loc = loc.Add(dir)
	}
	return locs, pipes
}

func getEnclosed(loopLocs []mathutil.Vec2, loopPipes []Pipe, grid [][]Traversable) ([]mathutil.Vec2, []Traversable) {
	castDirs := []mathutil.Vec2{mathutil.Vec2{1, 0}, mathutil.Vec2{-1, 0}, mathutil.Vec2{0, 1}, mathutil.Vec2{0, -1}}
	enclosedLocs := []mathutil.Vec2{}
	enclosedObjs := []Traversable{}
	height := len(grid)
	width := len(grid[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			loc := mathutil.Vec2{x, y}
			obj := grid[y][x]
			if slices.Contains(loopLocs, loc) {
				log.Printf("Skipping %v (%s) as part of loop", loc, obj.getSym())
				continue
			}
			misses := 0	
			for _, dir := range castDirs {
				ray := buildRay(loc, width-x, dir) 
				intersect := castRay(loopLocs, loopPipes, ray, grid)
				if intersect % 2 == 0 {
					misses ++	
				}
			}
			if misses < 3 {
				log.Printf("%v (%s) is enclosed!", loc, obj.getSym())
				enclosedLocs = append(enclosedLocs, loc)
				enclosedObjs = append(enclosedObjs, grid[y][x])
			}
		}
	}
	return enclosedLocs, enclosedObjs
}

func buildRay(start mathutil.Vec2, length int, dir mathutil.Vec2) []mathutil.Vec2 {
	ray := []mathutil.Vec2{}
	for i := 1; i < length; i++ {
		ray = append(ray, start.Add(dir.Mul(i)))
	}
	return ray
}

func castRay(contour []mathutil.Vec2, pipes []Pipe, ray []mathutil.Vec2, grid [][]Traversable) int {
	intersect := 0
	for _, loc := range ray {
		hitIdx := slices.Index(contour, loc)
		if hitIdx == -1 {
			continue
		}
		pipe := pipes[hitIdx]
		log.Print("Ray hit at ", loc, " for ", pipe.sym)
		intersect++
	}
	return intersect
}


func findInDirs(grid [][]Traversable, loc mathutil.Vec2) (mathutil.Vec2, mathutil.Vec2) {
	dirs := []mathutil.Vec2{mathutil.Vec2{1, 0}, mathutil.Vec2{0, 1}, mathutil.Vec2{-1, 0}, mathutil.Vec2{0, -1}}
	found := []mathutil.Vec2{}
	for _, dir := range dirs {
		x, y := loc.X+dir.X, loc.Y+dir.Y
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}
		if grid[y][x].isOpen(dir) {
			found = append(found, dir.Flip())
		}
	}
	return found[0], found[1]
}

func makeTraversable(lines []string) ([][]Traversable, mathutil.Vec2) {
	width := len(lines[0])
	height := len(lines)
	data := make([][]Traversable, height)
	for i := range data {
		data[i] = make([]Traversable, width)
	}
	var startLoc mathutil.Vec2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch sym := string(lines[y][x]); sym {
			case "|":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{0, -1}, sym}
			case "-":
				data[y][x] = Pipe{mathutil.Vec2{1, 0}, mathutil.Vec2{-1, 0}, sym}
			case "L":
				data[y][x] = Pipe{mathutil.Vec2{0, -1}, mathutil.Vec2{-1, 0}, sym}
			case "J":
				data[y][x] = Pipe{mathutil.Vec2{0, -1}, mathutil.Vec2{1, 0}, sym}
			case "7":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{1, 0}, sym}
			case "F":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{-1, 0}, sym}
			case ".":
				data[y][x] = Ground{sym}
			case "S":
				startLoc = mathutil.Vec2{x, y}
			}
		}
	}
	inA, inB := findInDirs(data, startLoc)
	data[startLoc.Y][startLoc.X] = Pipe{inA, inB, "S"}
	return data, startLoc
}
