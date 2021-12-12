package main

import (
	"fmt"
	"strings"
)

func explorePath(start string, connections map[string][]string, path []string, pathsCount int, 
	anySmallCaveVisitedTwice bool) int {
	path = append(path, start)

	if start == "end" {
		return pathsCount + 1
	}

	candidates := connections[start]
	for _, candidate := range candidates {
		if candidate == "start" {
			continue
		}
		smallCaveVisitedTwice := anySmallCaveVisitedTwice
		if candidate[0] >= 'a' && candidate[0] <= 'z' {
			smallCaveVisitsCount := 0
			for _, pathCell := range path {
				if pathCell == candidate {
					smallCaveVisitsCount++
				}
			}
			if smallCaveVisitsCount >= 1 {
				if smallCaveVisitedTwice {
					continue
				} else {
					smallCaveVisitedTwice = true
				}
			}
		} 
		pathsCount = explorePath(candidate, connections, path, pathsCount, smallCaveVisitedTwice)
	}
	return pathsCount
}

func main() {
	inputRows := []string{
		"start-kc",
		"pd-NV",
		"start-zw",
		"UI-pd",
		"HK-end",
		"UI-kc",
		"pd-ih",
		"ih-end",
		"start-UI",
		"kc-zw",
		"end-ks",
		"MF-mq",
		"HK-zw",
		"LF-ks",
		"HK-kc",
		"ih-HK",
		"kc-pd",
		"ks-pd",
		"MF-pd",
		"UI-zw",
		"ih-NV",
		"ks-HK",
		"MF-kc",
		"zw-NV",
		"NV-ks",
	}

	connections := make(map[string][]string, len(inputRows))

	for _, inputRow := range inputRows {
		parts := strings.Split(inputRow, "-")
		point1 := parts[0]
		point2 := parts[1]
		connections[point1] = append(connections[point1], point2)
		connections[point2] = append(connections[point2], point1)
	}

	pathsCount := explorePath("start", connections, nil, 0, false)

	fmt.Println(pathsCount, "paths")
}