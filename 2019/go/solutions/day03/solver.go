package day03

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 3
type Solver struct {
	Wires
	Panel         [][]int
	Intersections []utils.Point2D
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	lines := utils.TrimmedLines(fileContent)
	wires := Wires{}

	for _, line := range lines {
		wireInstructions := WireSegmentInstructions{}

		rawInstructions := strings.Split(line, ",")
		for _, rawInstruction := range rawInstructions {
			rawDirection, rawUnits := rune(rawInstruction[0]), rawInstruction[1:]

			units, err := strconv.Atoi(rawUnits)
			if err != nil {
				return fmt.Errorf("could not atoi units from value %v: %w", rawUnits, err)
			}

			direction, err := DirectionFromRune(rawDirection)
			if err != nil {
				return fmt.Errorf("error converting raw direction %v to Direction: %w", rawDirection, err)
			}

			instruction := WireSegmentInstruction{
				Direction: direction,
				Units:     units,
			}

			wireInstructions = append(wireInstructions, instruction)
		}

		wires = append(wires, Wire{SegmentsDefinition: wireInstructions})
	}

	s.Wires = wires
	return nil
}

func GetContainingRectangle(wires ...Wire) utils.Rectangle {
	minX, maxX, minY, maxY := 0, 0, 0, 0

	// For each wire
	for _, wire := range wires {
		x, y := 0, 0
		// For each instruction of this wire
		for _, wireInstruction := range wire.SegmentsDefinition {
			switch wireInstruction.Direction {
			case DOWN:
				y -= wireInstruction.Units
				if y < minY {
					minY = y
				}
			case UP:
				y += wireInstruction.Units
				if y > maxY {
					maxY = y
				}
			case LEFT:
				x -= wireInstruction.Units
				if x < minX {
					minX = x
				}
			case RIGHT:
				x += wireInstruction.Units
				if x > maxX {
					maxX = x
				}
			}
		}

	}

	return utils.Rectangle{
		Corner1: utils.Point2D{
			X: maxX,
			Y: maxY,
		},
		Corner2: utils.Point2D{
			X: minX,
			Y: minY,
		},
	}
}

func IntersectionSignalDelay(wire Wire, intersction utils.Point2D) int {
	steps := 0
	x, y := 0, 0

	// For each instruction of this wire
	for _, wireInstruction := range wire.SegmentsDefinition {
		units, dir := wireInstruction.Units, wireInstruction.Direction

		switch dir {

		case DOWN:
			for layed := 0; layed < units; layed++ {
				if x == intersction.X && y == intersction.Y {
					return steps
				}
				y--
				steps++
			}
		case UP:
			for layed := 0; layed < units; layed++ {
				if x == intersction.X && y == intersction.Y {
					return steps
				}
				y++
				steps++
			}
		case LEFT:
			for layed := 0; layed < units; layed++ {
				if x == intersction.X && y == intersction.Y {
					return steps
				}
				x--
				steps++
			}
		case RIGHT:
			for layed := 0; layed < units; layed++ {
				if x == intersction.X && y == intersction.Y {
					return steps
				}
				x++
				steps++
			}

		}
	}
	return steps
}

func LayWires(panel [][]int, centralPortPos utils.Point2D, wires ...Wire) {
	zx, zy := centralPortPos.X, centralPortPos.Y

	// For each wire
	for _, wire := range wires {
		x, y := 0, 0

		visitedPositions := utils.NewStringSet()

		// For each instruction of this wire
		for _, wireInstruction := range wire.SegmentsDefinition {
			units, dir := wireInstruction.Units, wireInstruction.Direction

			switch dir {

			case DOWN:
				for layed := 0; layed < units; layed++ {
					stringPos := fmt.Sprintf("%v,%v", x, y)
					if !visitedPositions.Has(stringPos) {
						panel[y+zy][x+zx]++
						visitedPositions.Add(stringPos)
					}
					y--
				}
			case UP:
				for layed := 0; layed < units; layed++ {
					stringPos := fmt.Sprintf("%v,%v", x, y)
					if !visitedPositions.Has(stringPos) {
						panel[y+zy][x+zx]++
						visitedPositions.Add(stringPos)
					}
					y++
				}
			case LEFT:
				for layed := 0; layed < units; layed++ {
					stringPos := fmt.Sprintf("%v,%v", x, y)
					if !visitedPositions.Has(stringPos) {
						panel[y+zy][x+zx]++
						visitedPositions.Add(stringPos)
					}
					x--
				}

			case RIGHT:
				for layed := 0; layed < units; layed++ {
					stringPos := fmt.Sprintf("%v,%v", x, y)
					if !visitedPositions.Has(stringPos) {
						panel[y+zy][x+zx]++
						visitedPositions.Add(stringPos)
					}
					x++
				}

			}
		}

	}

}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {

	// Get containing rectangle and set dimensions
	rect := GetContainingRectangle(s.Wires...)
	rows, columns := rect.AmplitudeY(), rect.AmplitudeX()
	panel := utils.MakeIntMatrix(rows+1, columns+1)

	// Get center offset
	bl := rect.BottomLeftCorner()
	center := utils.Point2D{
		X: -bl.X,
		Y: -bl.Y,
	}

	LayWires(panel, center, s.Wires...)
	s.Panel = panel

	var intersections []utils.Point2D

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			if panel[y][x] > 1 && y != center.Y && x != center.X {
				intersections = append(intersections, utils.Point2D{
					X: x - center.X,
					Y: y - center.Y,
				})
			}

		}
	}

	s.Intersections = intersections

	min := utils.ManhattanDistance(rect.BottomLeftCorner(), rect.TopRightCorner())
	for _, intersection := range intersections {
		dist := utils.ManhattanDistance(intersection, utils.Point2D{X: 0, Y: 0})
		if dist < min && intersection.X != 0 && intersection.Y != 0 {
			min = dist
		}
	}

	return strconv.Itoa(min), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {

	minSignal := 1_000_000

	for _, intersection := range s.Intersections {
		intersectionSignal := 0
		for _, wire := range s.Wires {
			intersectionSignal += IntersectionSignalDelay(wire, intersection)
		}
		if intersectionSignal < minSignal {
			minSignal = intersectionSignal
		}
	}
	return strconv.Itoa(minSignal), nil
}
