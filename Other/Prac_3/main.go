package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go rotateLeft.sherlockAndAnagrams
	if len(A) == 0 {
		return 0
	}

	for idx, value := range A {
		if idx+1 < len(A) {
			if A[idx+1]-value > 1 {
				return A[idx+1] - 1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(Solution([]int{}))
}
