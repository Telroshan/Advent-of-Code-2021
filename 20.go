package main

import (
	"fmt"
)

func enhanceImage(image []uint8, enhancementTable []uint8, imageWidth int, imageHeight int) (
	newImage []uint8, newImageWidth int, newImageHeight int) {
	newImageWidth = imageWidth + 2
	newImageHeight = imageHeight + 2
	newImage = make([]uint8, newImageWidth * newImageHeight)

	for y := 0; y < newImageHeight; y++ {
		for x := 0; x < newImageWidth; x++ {
			outputValue := 0

			bitIndex := 0
			for inputOffsetY := -1; inputOffsetY <= 1; inputOffsetY++  {
				inputY := y + inputOffsetY - 1
				for inputOffsetX := -1; inputOffsetX <= 1; inputOffsetX++  {
					inputX := x + inputOffsetX - 1
					inputValue := 0
					if inputX >= 0 && inputY >= 0 && inputX < imageWidth && inputY < imageHeight {
						if image[inputY * imageWidth + inputX] > 0 {
							inputValue = 1
						}
					}
					
					outputValue |= (inputValue << (8 - bitIndex))

					bitIndex++
				}
			}

			newImage[y * newImageWidth + x] = enhancementTable[outputValue]
		}
	}

	return
}

func printImage(image []uint8, imageWidth int, imageHeight int) {
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			if image[y * imageWidth + x] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func main() {
	inputRows := []string{
		"######.#..##..######.....#...#.#...#.######.#.#...#..#..#..###.#####.#.####...#.#.#...#.#.#...#####..###.....#.#.....#.#.#..###..#.#.#####..#.....####.##.##..#..##.#.##.##..##.#####.####.#.#....#...#...#...#.#########..#..#####..#.#.###....#.##.###...##.#..#....##.###...#..##.#..#...#.#.#####.####...#.##..##..#.#######...#..##.##.....#..#..#.###...###.######..##.#..##.######....#.##.##...###.##..#.#.#.#########....####.####.#.#...#.#.#..#.#.##.#...#.#..#######..###..##.#..####.###...#.#.#.##.#####.##.###.#.",
		"",
		".##.##...#####.#######.##.##.#####..#.#.#.##.....##...#..#..###.#....###...##.....###......#..##...#",
		"#..#...#.##..#.#..#....#...#...#....####..#.....#.#...##.#..##.##.##...#.#..#..####...#....##.##.##.",
		"...#.#..#..###...##..#.##########...##...#.#.#.#.#####..###..###..##.#..#####..###......##.#.##...##",
		"#..###.##.##...#.#..#.....#.##.#..#...#..####.#.#.##.###....#.#..#..##.#....#.#.#.###...###.#.####.#",
		"...###.....###...##...#..######.###...##.#.##.###.#.#....###.##.###.###.###....######..##.#.#...#..#",
		"###...#.##.#.###..#..#.####.#...#..#.#.#..#.#..#....###..##..####.##.##....#..##.#.##..#.#..#.#.##..",
		".###.##.#.....#.#.##.##.####.#..#.#..###.##...###..##..##.#.#...#..#.#####...#.####.#....######.#..#",
		"..##.........##.##......##..#.#.#....#..####..#.#..###.##..#.##..####.#.#.##.#.#.##.....##....#.####",
		"#..###..#...###.##.#.#....#.#....###....###...##.#..####.#####.#.##.#...#.#.#.##..#....#..#.#...##.#",
		"###.#..#...#.##.#.#.#..#.####.###...##.####.#.#.........#.##..####.###..#.#...###..#.##..##...####..",
		"#....###.##.##...###...##.###...#..##....#.##.#.##.#.##..#..##..#..#..#...#.##..##.###.#..##.#..###.",
		".#.#......##..#..#....#.##.#..##.##....#....#..###..##.##..###..#..#.#.##.#.##..#...#.#....#.#...#.#",
		"#.#..#....#.##.##..#.###..##..####.##...##.#.#.###.#.#..#...#.##..#..#.###.....#.##.#####..#......#.",
		"##.##...#.....##.....###.#..##...#.#...#.##..#........#######...#.###.#.#.##..##.#...###..#...#...#.",
		"....###...###.#.#.#.###..#..####..######....###....####..##..#..#.#.###.#.#.#####...#..#.#.####...##",
		"....#######.#....##...#####...###...#...#...#.#...#.###.#.####.#...#.#.#.###....#..#....#.#.#..##..#",
		"#.####.###.#..#.#.#..#####.......##....####.##.####...#..###..#......#..#.#...#....##.....#.........",
		"#..###.#..##.##..###.#.#.#.###.#....#.####.#.###.##.##..#...#.#.#..#..###..#..##......#.####..#.#.##",
		"###.#..##.###....##.#####.###.....#....#####.#..#.#.#.##..#...##...###.##.##...###.##..#..#...#.#..#",
		"#.#...#..#..#.###.#..#..##.#.#.....#.#.#....#.##...##.......####..#...#..#.......###..#.#.#....#..##",
		".##...##..#..#..####.##.......#...#.###..#..##.#.#.#.##..#..##.###.##.#.#.#..##.##.##..#..###.#..#.#",
		"#...#.###.#......#...#.......##..##...##...#.#....##.#.##.##....##.######..##.##.####.#####....#.#..",
		"#.##.####....##....#..#.#..#.####.#.####.###..##..#.....##.###.#...#...#..#..##.##.###...###.###.#.#",
		"....#.##.##.###.#...#....#...###.#.#......#..#.###.#.#.#.#...#.#..###..#.##.#...##.######..#..####..",
		"####...#.#.###.#...#.......##..#.#######...##...##..##......###.#####.#.#.#.###.#..#.#..##.##..###.#",
		"#.#......#.#.####..##.##...#..######..#.#..#...#..#.####.#.##.#####.#..##..##.#.###.#.#.##.#.#.##.#.",
		"..#.#..#.#...###.#..#.#.##.#.#....###.#..###..##..#######.#.##.####.....#.####.#.##.#...#####......#",
		"#######...##.#.#.###..###.#..##..##...####..###.#.##.##.##...#.##.##.##....#.#..#.###...###.#.#.#...",
		".##.##..####..#.###########....#...#.#..##......##......#.#..###.##....#..#..#.#..###########..##.##",
		"#....##...###..#.#.###.##......###.#.###.######.#..#.###.#.#.##..#....###.###.#.....#..#..#.#####..#",
		".######..#.###...###.#..##..#.####.###...###.#.#.####.########.#...#.#.##....####.###.#.##.###...#.#",
		"###..##.##.....#...#.##...#.###...#.#.####..#.##..##..#....#..##.#######.##...#####...##.#...##.##..",
		"###.#.#..###..##.#...####.#..###....###..###..###.###..#.##...#.....#...###.##.#.###.#####.#.##.#.#.",
		"###..######.....#..#.....##.##.##.##...#.#..####....###.#.##...####.#....##.......##.#..###......#.#",
		"##..#.###.##.######.#.##.#...#...###......##.###....##...#.#####.#.##.....#.#..#...####...###...##..",
		"...###..###..##.##..####...##....##.#...#.##..###.###.####...##.#.#.##.#....#.#.###.#..#..#..#.#.##.",
		"..#.###.#.#...#.###..##..##....#.#...#.##.####....#.#.#.#..###.#.#....###.##.#.#...#...#...#..##.#..",
		"###...#.##..##.#..###.#.###.#..#.#.##.#.###...###.#...##.#.......#.#.###....####.....###.##.##.#.#..",
		".#.####..##.#.#..#######..#..###.#########.#...##..#..##.#.....#.#.##.#.....#.#######..###.....##..#",
		"####.##....#..##.....#...#.#....##.#...####..###.###..#......###.###....##.#.#....###.###.#####..#..",
		".#.#...##.##.#...#.#...##..###.##..####.###..#....#######..#..##...#####..##...#.#.#######.####.#...",
		".#.###.######...##.##.#.##...####.#.......#.####.##.#.##.##.####.#...#.#..##....#####.##.#.##..#.#..",
		"..#.#..###.....##....#.###....####.#.##.######..###.##..############.#...#...###.##.#..##.#..####..#",
		".#.....#.###.#.##....#...#.#####.#.#.##..#...###..##....#..#......#####.#.###..#..####...#####......",
		".##..#..........##.#.######....##.#.#...###.........#.##..##.......#..#.##......###.....##....#....#",
		"#...#..#.#.#.#.#.####.#..#...#....##.#.#.#....##..#####......##.###.##..#..##...#.##...#....#.##.#..",
		"...#..#.##.####.###.####....##...#......###.###.##....#...##.##.#..#.##.#.....#.##.#.##.##.#.###...#",
		"..#..##.##.#.#####...#..###.#..###..#.####...#.#.##.#...#...#.#.#.###.#.#..##.##..#.##..#....#..###.",
		"..#.#.###...##.#...#.#.#####...###.##...#..###.....#...##.#....##...###.##.####..#..##....#####.##.#",
		"..#.#...#.###...##.##....##.####.....##.##....###..#...#.###..##.#.#.#.##..##...##....#.###.####.#.#",
		"#...#####.##.#.#.#.#...#..###..#.##..##....#..#.###...#.####.###.##.#.##.##......####.....#####..###",
		".#.#.....#.######.#..#.##..###..#.###.###.##..#.##.....#......#..##.#...###.###..#.#..#.###...#.####",
		"####.##.##.###..#.##.##.#.###.#...##...#..#.###.....#..##..###..#.#..#...#.##.#..#.##.#..#..#.#..#.#",
		"#####.##.####.#######.#.##..##..#.#####..#..#.###..###..#..#.##...#..##.#####...###.#..##.#.###.#..#",
		"...##..##.#.#.###..#.##.#.###....#...###.#####.#.#.#.#......#.###..#..##..#........#.#..#.##..#....#",
		"#.#..#####.##.###...#...###..#.###.......#.#..#..#.#...###....####.###........#.###.####..####.#..#.",
		"........#.##....#.##...#..##...##.####.##.#.##..##......###.#..#.#####..##.#.......##....##..##..#.#",
		"..####.##...###..#####...#..####..###.##.#.#.#.#.....##.##...#.#....###..#..#.##....#...#.####...##.",
		"..##..#.##....##.####.#.#..####.....#.###..#...##..##.#......##...###.....##.....#....##...##.##..##",
		".#.##...#####..##.##.#..#...#####....#.....#.##.#.##..#...###########...###..#.##.###....#.##.#####.",
		"####......####.#####..###...#.####..#.#..###..#...##..##..#.#.#####.###.##.##..#####.##.#.###.##..#.",
		"##.#..#.####.##..#..#.###.....###.#####.#.#####......#...##..#....##.#.##..###.#####..#.##.#..#..###",
		".##..####..###....#...#..#.###.#.##..#..#.###..###.##.##...#...##.#..#...#.###..#.###.#.##..##.##.##",
		"#.##..##.#....###.###...#.#.###..#.######..##.#....###.#.###.####..#...#.#..#.#.##.###...##...#.#...",
		"#.#.##.#..#.##....###..###...#...#.#..##.#..##.#..#.##..#.#..####.#....###.#..##..#...#.###.###..#.#",
		"##.##..#####..##.##.##.##..#.....###..#..###..#.##.##.##.#..##.#..#..###.####..#.#####.##.#####.#.##",
		"...###.#...#...###......###.#....##..##.##..########..###....###.####.#..###.######.##.......####.##",
		"##.#.#..#.##.##..#...#...#....#....#....##.#..##.#......#.#.#.#...#.#.####.#.##..#..#.##.##.####...#",
		"##.##...#.#.#.##..#.#....##...###.#####..###..#.###.###.#..###......##.###....##.#.##.####....#####.",
		".#.#.#####.#......#.##.#.###.#...#.##..#.#....###.##...#...#..##.####..####..###...#..##......###...",
		"###.....##...#######..#.####.....#...#.#...#..#.###.##..#.#....#.#####..##.#.#.#####..#.#.####.#.##.",
		"..#..##..#.#..#...#.#.....##.#.#..##..#.#...#.##.....#...###..##..#.#.#..####..#..#.#.###...#.##.#.#",
		"#...##..#.######.##.#.###.#....#..#.#....#...########.....#####.####....#.##.##..#.#.#...#..##..###.",
		".#####..####.#..#..#.##.#...#..##.#..##...#....##.#.###......#.......#.....###.....##.####...##..###",
		"##.#...#......###.#.....#..#######.##..#.#.#####..#....##..#..##.#...#.##..#...#.#...#.#.##.#.#####.",
		"#.###..#.####..##..#######....#######..##...###..#.#..#....#.####....#...####.#..###.####...##.#.##.",
		".#..##..#.#.#.#######...#.#.#.#..#.....###.##..###...#.#.##.#..#.##...#..##..#...#.#...###...####...",
		"...#.#.#.....#.######..##..#.###.###.#.#.###..#....##...###......#.##..#####...###...#...##.#...##..",
		"#.#.......#.##......#...######.....#.###..#####..#.###.##..##...###.##..######...##.#..###.##.#..#.#",
		"....#..##.#.....###.###.#....###..####....#.....###..#.....##.##.##.#..#.#.##...#.##..#.#....###...#",
		"####....#.###..########.##......###....#####.#.#..###.#.#.##.######.#...##...#.###.##.#.....####..##",
		"#####.####....#....###...#.#.#.###.#####..##...#...#.##..#######...#......#...#..####.##......#####.",
		"..##.##.#..#.#.##..##.##..#..#....#.....#######...#..#......#.#.##...#.#.#....###...##..##.#...#.##.",
		".##.##.#....##.###....#.####....####..##....######.....####.####.#..#.....#####.##.#.##.######....#.",
		"..#...#......##...#..#.#...#.#####..###..#.#..####....#...##.....#.#.#.##.##..#.#.###.##..#.####.#..",
		".#..#.##.#...#....####.##..##..#####..#..#.#####...##.#.#..####.#...##....#..#####.##.##..##.#..#..#",
		".##..#.###...#.#..###.###.#......######..#.#.####...#...#..##.########.##..#.###......#.##.#...#.#.#",
		"...#.#..##..#....#..#####.#..######.#..#.#.#...##.######..##..####......#.##.##.#..#..#.##.##...#.#.",
		"#..#.###.####.#.....#...#.....#..##.......#...#..###.###.#.....#.##.#..#.###....##..#...##.#....#...",
		".##...#..##..##.##...##.##..#.#.#..#.#.#####.####....#....#.#..#..#.#..#...#..#....#.#.#.######....#",
		"#..##.##..#...#....#####....#####...#..#.##.##..##.#..#####.#..####.##.#.#####.#..##.#......##.#..##",
		"#...##.#...#..###....####.##.#.#..#.#.#.#.#..#.#..#..##.#.#.#####...##.##.##.##.##.###...##..#...#.#",
		"#.#..###.....#..####.#.###.######...#######.####..#..##..#.....##.###..#.#.##.###.#.#..##.#.###..###",
		"....###..###.###....####....##..##...#..#...##.#...###..#..#.##...#......###.####...#...#..#.....#.#",
		"#.#..##.##.....####.#..##...#.##.#..#.#.....##.####.#.##.##..#.....####.###.##..##...##...##.#..##.#",
		"######..#.###.##..#......#.###.#.#.....##.###...##..##.###....##....##.##.#.#.#...#..#..#..###..#..#",
		"..#...#...######.##.#.##..#.#.#.#.#.#.#####.####.######.####.##.#.#.#.#....#..#....####.#.##..#.#.##",
		"####..#.##.#..###...##...###.###.####...#.##.####.###.###.##..####.##.#....#..###..#..##.##.#.....#.",
		".##.##.##.....#######.#..##.#...#.#####.#....##..#...#..##...##...#..#.#..#...#.#..#.###.###.##.###.",
		"...##.#.#.##..#..#..#.#.....##..#.##..##.#..###..#...###......#.###..##..#####..##.###.##....#.##.#.",
	}

	enhancementTable := make([]uint8, 512)
	for i, char := range inputRows[0] {
		if char == '#' {
			enhancementTable[i] = 1
		} else {
			enhancementTable[i] = 0
		}
	}

	imageStartLine := 2
	imageHeight := len(inputRows) - imageStartLine
	imageWidth := len(inputRows[imageStartLine])
	image := make([]uint8, imageWidth*imageHeight)
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			char := inputRows[imageStartLine+y][x]
			if char == '#' {
				image[y*imageWidth+x] = 1
			} else {
				image[y*imageWidth+x] = 0
			}
		}
	}
	
	fmt.Println("STEP 0")
	printImage(image, imageWidth, imageHeight)

	for i := 1; i <= 2; i++ {
		fmt.Println("\n\n---------\nSTEP", i)
		image, imageWidth, imageHeight = enhanceImage(image, enhancementTable, imageWidth, imageHeight)
		
		printImage(image, imageWidth, imageHeight)
	}

	litPixelsCount := 0
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			if image[y * imageWidth + x] > 0 {
				litPixelsCount++
			}
		}
	}

	fmt.Println(litPixelsCount, "lit pixels")
}
