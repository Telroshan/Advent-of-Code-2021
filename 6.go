package main

import "fmt"

func main() {
	values := []int{1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 1, 4, 2, 4, 1, 1, 1, 1, 1, 5, 4, 1, 1, 2, 1, 1, 1, 1, 4, 1, 1, 1, 4, 4, 1, 1, 1, 1, 1, 1, 1, 2, 4, 1, 3, 1, 1, 2, 1, 2, 1, 1, 4, 1, 1, 1, 4, 3, 1, 3, 1, 5, 1, 1, 3, 4, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 2, 5, 5, 3, 2, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 5, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 2, 4, 1, 5, 5, 1, 1, 5, 3, 4, 4, 4, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 5, 3, 1, 4, 1, 1, 2, 2, 1, 2, 2, 5, 1, 1, 1, 2, 1, 1, 1, 1, 3, 4, 5, 1, 2, 1, 1, 1, 1, 1, 5, 2, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 5, 1, 4, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 5, 4, 5, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 3, 1, 1, 1, 3, 1, 4, 2, 1, 5, 1, 3, 5, 5, 2, 1, 3, 1, 1, 1, 1, 1, 3, 1, 3, 1, 1, 2, 4, 3, 1, 4, 2, 2, 1, 1, 1, 1, 1, 1, 1, 5, 2, 1, 1, 1, 2}
	// fmt.Println(len(values), "fishes at start")
	// for day := 0; day < 80; day++ {
	// 	for i := range values {
	// 		values[i]--
	// 		if values[i] < 0 {
	// 			values[i] = 6
	// 			values = append(values, 8)
	// 		}
	// 	}
	// 	// fmt.Println(len(values), "fishes")
	// }
	// fmt.Println(len(values), "fishes")

	// type Group struct {
	// 	Lifetime int
	// 	Count int
	// }
	// initialGroups := make([]Group, 7)
	// for _, value := range values {
	// 	group := initialGroups[value]
	// 	group.Lifetime = value
	// 	group.Count++
	// 	initialGroups[value] = group
	// }
	// fmt.Println(initialGroups)

	fishCounts := make([]int, 9)
	for _, value := range values {
		fishCounts[value]++
	}

	resetCount := 0
	for day := 0; day < 256; day++ {
		for i := 0; i < len(fishCounts); i++ {
			if i > 0 {
				fishCounts[i-1] = fishCounts[i]
			} else {
				resetCount = fishCounts[0]
			}
		}
		fishCounts[6] += resetCount
		fishCounts[8] = resetCount
		resetCount = 0
	}

	total := 0
	for _, count := range fishCounts {
		total += count
	}
	fmt.Println(total, "fish")
}
