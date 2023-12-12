package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/EliCDavis/vector/vector2"
	"github.com/fogleman/gg"
	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

type Ground struct{ sym string }

type Pipe struct {
	loc, inA, inB vector2.Vector[int]
	sym           string
}

type Traversable interface {
	isOpen(vector2.Vector[int]) bool
	getSym() string
}

func (g Ground) isOpen(inDir vector2.Vector[int]) bool { return false }
func (g Ground) getSym() string                        { return g.sym }

func (p Pipe) isOpen(inDir vector2.Vector[int]) bool {
	return inDir == p.inA || inDir == p.inB
}

func (p Pipe) getOutDir(inDir vector2.Vector[int]) vector2.Vector[int] {
	var out vector2.Vector[int]
	if inDir == p.inA {
		out = p.inB.Flip()
	} else if inDir == p.inB {
		out = p.inA.Flip()
	} else {
		log.Fatalln("Pipe ", p, " not open to dir ", inDir)
	}
	return out
}

func (p Pipe) getSym() string { return p.sym }

func main() {
	lines := []string{}
	file, err := os.Open("../input/10-test-input-3.txt")
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
	locs, loop := walkLoop(grid, start)
	fmt.Println("Part 1 answer :", len(loop)/2)

	loopSegs := buildLineSegs(loop)
	width, height := float64(len(grid[0])), float64(len(grid))
	enclosed, locs := countEnclosed(loopSegs, locs, grid)
	Draw(loopSegs, locs, width, height, 16.0, "./out.png")
	fmt.Println("Part 2 answer :", enclosed)
}

func walkLoop(grid [][]Traversable, start vector2.Vector[int]) ([]vector2.Vector[int], []Pipe) {
	startPipe := grid[start.Y()][start.X()].(Pipe)
	dir := startPipe.inA.Flip()
	pipes := []Pipe{startPipe}
	locs := []vector2.Vector[int]{start}
	loc := start.Add(dir)
	for loc != start {
		locs = append(locs, loc)
		pipe := grid[loc.Y()][loc.X()].(Pipe)
		pipes = append(pipes, pipe)
		dir = pipe.getOutDir(dir)
		loc = loc.Add(dir)
	}
	return locs, pipes
}

func buildLineSegs(pipes []Pipe) []mathutil.LineSeg {
	segs := []mathutil.LineSeg{}
	for _, pipe := range pipes {
		offsetA := pipe.inA.ToFloat64().Flip().Scale(0.5)
		a := pipe.loc.ToFloat64().Add(offsetA)
		offsetB := pipe.inB.ToFloat64().Flip().Scale(0.5)
		b := pipe.loc.ToFloat64().Add(offsetB)
		segs = append(segs, mathutil.LineSeg{a, b})
	}
	return segs
}

func countEnclosed(loop []mathutil.LineSeg, loopLocs []vector2.Vector[int], grid [][]Traversable) (int, []vector2.Vector[int]) {
	height := len(grid)
	width := len(grid[0])
	rayOffset := vector2.New[float64](0.875, 0.5145).Scale(float64(height + width))

	enclosed := 0
	locs := []vector2.Vector[int]{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			loc := vector2.New[int](x, y)
			obj := grid[y][x]
			if slices.Contains(loopLocs, loc) {
				log.Printf("Skipping %v (%s) as part of loop", loc, obj.getSym())
				continue
			}
			floc := loc.ToFloat64()
			rayEnd := floc.Add(rayOffset)
			ray := mathutil.LineSeg{floc, rayEnd}
			intersect := mathutil.CountIntersections(ray, loop)
			if intersect%2 == 1 {
				enclosed++
				locs = append(locs, loc)
			}
		}
	}
	return enclosed, locs
}

func findInDirs(grid [][]Traversable, loc vector2.Vector[int]) (vector2.Vector[int], vector2.Vector[int]) {
	dirs := []vector2.Vector[int]{vector2.New[int](1, 0), vector2.New[int](0, 1), vector2.New[int](-1, 0), vector2.New[int](0, -1)}
	found := []vector2.Vector[int]{}
	for _, dir := range dirs {
		x, y := loc.X()+dir.X(), loc.Y()+dir.Y()
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}
		if grid[y][x].isOpen(dir) {
			found = append(found, dir.Flip())
		}
	}
	return found[0], found[1]
}

func makeTraversable(lines []string) ([][]Traversable, vector2.Vector[int]) {
	width := len(lines[0])
	height := len(lines)
	data := make([][]Traversable, height)
	for i := range data {
		data[i] = make([]Traversable, width)
	}
	var startLoc vector2.Vector[int]

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			loc := vector2.New[int](x, y)
			switch sym := string(lines[y][x]); sym {
			case "|":
				data[y][x] = Pipe{loc, vector2.New[int](0, 1), vector2.New[int](0, -1), sym}
			case "-":
				data[y][x] = Pipe{loc, vector2.New[int](1, 0), vector2.New[int](-1, 0), sym}
			case "L":
				data[y][x] = Pipe{loc, vector2.New[int](0, -1), vector2.New[int](-1, 0), sym}
			case "J":
				data[y][x] = Pipe{loc, vector2.New[int](0, -1), vector2.New[int](1, 0), sym}
			case "7":
				data[y][x] = Pipe{loc, vector2.New[int](0, 1), vector2.New[int](1, 0), sym}
			case "F":
				data[y][x] = Pipe{loc, vector2.New[int](0, 1), vector2.New[int](-1, 0), sym}
			case ".":
				data[y][x] = Ground{sym}
			case "S":
				startLoc = vector2.New[int](x, y)
			}
		}
	}
	inA, inB := findInDirs(data, startLoc)
	data[startLoc.Y()][startLoc.X()] = Pipe{startLoc, inA, inB, "S"}
	return data, startLoc
}

func Draw(segs []mathutil.LineSeg, points []vector2.Vector[int], width, height, scale float64, path string) {
	fullWidth, fullHeight := width*scale, height*scale
	dc := gg.NewContext(int(fullWidth), int(fullHeight))

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Flip on Y axis
	dc.ScaleAbout(1, -1, fullWidth/2.0, fullHeight/2.0)

	dc.SetRGB(0, 0, 0)
	offset := vector2.New[float64](0.5, 0.5).Scale(scale)
	for _, seg := range segs {
		a := seg.A.Scale(scale).Add(offset)
		b := seg.B.Scale(scale).Add(offset)
		dc.DrawLine(a.X(), a.Y(), b.X(), b.Y())
		dc.Stroke()
	}

	dc.SetRGB(0.9, 0.1, 0.1)
	for _, point := range points {
		loc := point.ToFloat64().Scale(scale).Add(offset)
		dc.DrawPoint(loc.X(), loc.Y(), 2)
		dc.Fill()
	}

	// ray vector
	dc.SetRGB(0.1, 0.1, 0.9)
	ray := vector2.New[float64](0.875, 0.5145).Scale(float64(height + width)).Scale(scale)
	dc.DrawLine(0.0, 0.0, ray.X(), ray.Y())
	dc.Stroke()

	// Grid
	dc.SetRGB(0.0, 0.5, 0.5)
	for x := offset.X(); x <= fullWidth; x += scale {
		for y := offset.Y(); y <= fullHeight; y += scale {
			dc.DrawPoint(x, y, 1)
			dc.Fill()
		}
	}

	err := dc.SavePNG(path)
	if err != nil {
		log.Fatal(err)
	}
}
