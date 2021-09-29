package main

import "fmt"

func sliceContains(src []int, value int) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

func Solution(ranks []int) int {
	// write your code in Go rotateLeft.sherlockAndAnagrams
	var (
		result int
	)
	result = 0
	for _, value := range ranks {
		if bool := sliceContains(ranks, value+1); bool {
			result += 1
		}
	}

	return result
}

func main() {
	//fmt.Println(Solution([]int{minimumSwaps,sherlockAndAnagrams,minimumSwaps,0,minimumBribes,minimumBribes,minimumSwaps,0,0}))
	fmt.Println(Solution([]int{4, 2, 0}))
	fmt.Println(Solution([]int{4, 4, 3, 3, 1, 0}))
}
