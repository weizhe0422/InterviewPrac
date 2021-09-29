package main

import (
	"fmt"
)

func Solution(A int, B int) string {
	// write your code in Go rotateLeft.sherlockAndAnagrams
	var (
		resStr          []byte
		totalCnt        int
		mainCha, elsCha byte
		mainCnt, elsCnt int
	)

	if A > B {
		mainCha = 'A'
		elsCha = 'B'
		mainCnt = A
		elsCnt = B
	} else {
		mainCha = 'B'
		elsCha = 'A'
		mainCnt = B
		elsCnt = A
	}
	totalCnt = mainCnt + elsCnt

	resStr = make([]byte, totalCnt)
	for totalCnt > 0 {
		if totalCnt%3 != 0 {
			if mainCnt > 0 {
				resStr[totalCnt-1] = mainCha
				mainCnt -= 1
			} else {
				resStr[totalCnt-1] = elsCha
				elsCnt -= 1
			}
		} else {
			if elsCnt > 0 {
				resStr[totalCnt-1] = elsCha
				elsCnt -= 1
			} else {
				resStr[totalCnt-1] = mainCha
				mainCnt -= 1
			}
		}
		totalCnt -= 1
	}

	return string(resStr)
}

func main() {

	fmt.Println(Solution(5, 3))

}
