package main

import "fmt"

func incrementCell(grid []uint8, gridWidth int, gridHeight int, x int, y int) {
	index := y*gridWidth + x
	grid[index]++
	if grid[index] != 10 {
		return
	}
	if x > 0 {
		incrementCell(grid, gridWidth, gridHeight, x-1, y)
		if y > 0 {
			incrementCell(grid, gridWidth, gridHeight, x-1, y-1)
		}
		if y < gridHeight-1 {
			incrementCell(grid, gridWidth, gridHeight, x-1, y+1)
		}
	}
	if x < gridWidth-1 {
		incrementCell(grid, gridWidth, gridHeight, x+1, y)
		if y > 0 {
			incrementCell(grid, gridWidth, gridHeight, x+1, y-1)
		}
		if y < gridHeight-1 {
			incrementCell(grid, gridWidth, gridHeight, x+1, y+1)
		}
	}
	if y > 0 {
		incrementCell(grid, gridWidth, gridHeight, x, y-1)
	}
	if y < gridHeight-1 {
		incrementCell(grid, gridWidth, gridHeight, x, y+1)
	}
}

func main() {
	inputRows := []string{
		"5723573158",
		"3154748563",
		"4783514878",
		"3848142375",
		"3637724151",
		"8583172484",
		"7747444184",
		"1613367882",
		"6228614227",
		"4732225334",
	}

	gridWidth := len(inputRows[0])
	gridHeight := len(inputRows)
	grid := make([]uint8, gridWidth*gridHeight)

	for y, inputRow := range inputRows {
		for x, char := range inputRow {
			grid[y*gridWidth+x] = uint8(char - '0')
		}
	}

	flashesCount := 0
	allFlashed := false
	for step := 1; step <= 100 || !allFlashed; step++ {
		for y := 0; y < gridHeight; y++ {
			for x := 0; x < gridWidth; x++ {
				incrementCell(grid, gridWidth, gridHeight, x, y)
			}
		}

		flashesDuringThisStep := 0
		for y := 0; y < gridHeight; y++ {
			for x := 0; x < gridWidth; x++ {
				if grid[y*gridWidth+x] > 9 {
					grid[y*gridWidth+x] = 0
					if step <= 100 {
						flashesCount++
					}
					flashesDuringThisStep++
				}
			}
		}

		if !allFlashed && flashesDuringThisStep == gridWidth * gridHeight {
			fmt.Println("All flash simultaneously at step", step)
			allFlashed = true
		}
	}

	fmt.Println(flashesCount, "flashes in 100 steps")
}
