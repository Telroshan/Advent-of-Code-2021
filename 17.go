package main

import (
	"fmt"
	"strings"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	input := "target area: x=288..330, y=-96..-50"
	const prefix = "target area: x="
	parts := strings.Split(input[len(prefix):], ", y=")
	xParts := strings.Split(parts[0], "..")
	yParts := strings.Split(parts[1], "..")
	var minX int
	var maxX int
	var minY int
	var maxY int
	minX, _ = strconv.Atoi(xParts[0])
	maxX, _ = strconv.Atoi(xParts[1])
	minY, _ = strconv.Atoi(yParts[0])
	maxY, _ = strconv.Atoi(yParts[1])

	fmt.Println(minX,maxX,minY,maxY)

	highestReachedY := 0
	bestStartVelocityY := 0
	for yStartVelocity := abs(minY); yStartVelocity > 0; yStartVelocity-- {
		y := 0
		landsInTarget := false
		highestY := 0
		yVelocity := yStartVelocity
		for y > minY {
			y += yVelocity
			yVelocity--
			if y > highestY {
				highestY = y
				landsInTarget = false
			}
			if y <= maxY && y >= minY {
				landsInTarget = true
			}
		}
		if landsInTarget && highestY > highestReachedY {
			highestReachedY = highestY
			bestStartVelocityY = yStartVelocity
		}
	}
	fmt.Println("Best start :", bestStartVelocityY,"reaches :",highestReachedY)

	possibleVelocitiesCount := 0
	for xStartVelocity := 0; xStartVelocity <= maxX; xStartVelocity++ {
		for yStartVelocity := minY; yStartVelocity <= abs(minY); yStartVelocity++ {
			x := 0
			y := 0
			yVelocity := yStartVelocity
			xVelocity := xStartVelocity
			for y > minY {
				y += yVelocity
				x += xVelocity
				yVelocity--
				if xVelocity > 0 {
					xVelocity--
				}
				if y <= maxY && y >= minY && x <= maxX && x >= minX {
					possibleVelocitiesCount++
					break
				}
			}
		}
	}
	fmt.Println(possibleVelocitiesCount,"possible start velocities")
}