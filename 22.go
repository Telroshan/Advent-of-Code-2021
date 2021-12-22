package main

import (
	"fmt"
	"strconv"
	"strings"
	_ "embed"
)

//go:embed 22.txt
var fileContent string

func main() {
	inputRows := strings.Split(fileContent, "\r\n")

	type Instruction struct {
		xMin   int
		xMax   int
		yMin   int
		yMax   int
		zMin   int
		zMax   int
		turnOn bool
	}

	// instructions := make([]Instruction, len(inputRows))
	var instructions []Instruction

	var minX int = 1 << 62
	var maxX int
	var minY int = 1 << 62
	var maxY int
	var minZ int = 1 << 62
	var maxZ int

	for _, inputRow := range inputRows {
		var instruction Instruction
		instructionSplit := strings.Split(inputRow, " ")
		instruction.turnOn = instructionSplit[0] == "on"
		coordinatesSplit := strings.Split(instructionSplit[1], ",")
		for _, split := range coordinatesSplit {
			targetSplit := strings.Split(split, "=")
			valuesSplit := strings.Split(targetSplit[1], "..")
			min, err := strconv.Atoi(valuesSplit[0])
			if err != nil {
				panic(err)
			}
			max, err := strconv.Atoi(valuesSplit[1])
			if err != nil {
				panic(err)
			}
			switch targetSplit[0] {
			case "x":
				instruction.xMin = min
				instruction.xMax = max
			case "y":
				instruction.yMin = min
				instruction.yMax = max
			case "z":
				instruction.zMin = min
				instruction.zMax = max
			}
		}
		if instruction.xMin >= -50 && instruction.xMax <= 50 &&
			instruction.yMin >= -50 && instruction.yMax <= 50 && 
			instruction.zMin >= -50 && instruction.zMax <= 50 {
			// instructions[i] = instruction
			instructions = append(instructions, instruction)
			
			if instruction.xMin < minX {
				minX = instruction.xMin
			}
			if instruction.xMax > maxX {
				maxX = instruction.xMax
			}
			if  instruction.yMin < minY {
				minY = instruction.yMin
			}
			if instruction.yMax > maxY {
				maxY = instruction.yMax
			}
			if instruction.zMin < minZ {
				minZ = instruction.zMin
			}
			if instruction.zMax > maxZ {
				maxZ = instruction.zMax
			}
		}
	}

	gridWidth := maxX - minX + 1
	gridHeight := maxY - minY + 1
	gridDepth := maxZ - minZ + 1

	fmt.Println(instructions)

	fmt.Println(gridWidth,"*",gridHeight,"*",gridDepth,"\tx:",minX,"=>",maxX," y:", minY,"=>",maxY,
		" z:",minZ,"=>",maxZ)
	grid := make([]bool, gridWidth*gridHeight*gridDepth)
	turnedOnCubesCount := 0
	for _, instruction := range instructions {
		// fmt.Println(i+1,":",instruction)
		for z := 0; z <= instruction.zMax-instruction.zMin; z++ {
			for y := 0; y <= instruction.yMax-instruction.yMin; y++ {
				for x := 0; x <= instruction.xMax-instruction.xMin; x++ {
					coord := (z+instruction.zMin-minZ)*gridHeight*gridWidth +
						(y+instruction.yMin-minY)*gridWidth +
						(x + instruction.xMin - minX)
					// fmt.Println("\t",  x +instruction.xMin, y + instruction.yMin,
					// z + instruction.zMin,"\t",x,y,z ,"=>",coord,)
					if grid[coord] != instruction.turnOn {
						if instruction.turnOn {
							turnedOnCubesCount++
							// fmt.Println("\t\t+1")
						} else {
							turnedOnCubesCount--
							// fmt.Println("\t\t-1")
						}
					}
					grid[coord] = instruction.turnOn
				}
			}
		}
	}

	fmt.Println(turnedOnCubesCount, "turned on cubes")
}
