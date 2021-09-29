package main

import (
	"fmt"
	"log"
)

func count(input []int) (result int) {
	var (
		freq        map[int]int
		tgtCnt      int
		tmpSlice    []int
		resultSlice map[int]int
		sum         int
	)

	freq = make(map[int]int)
	resultSlice = make(map[int]int)
	tgtCnt = len(input) / 2

	for _, value := range input {
		freq[value] = freq[value] + 1
	}

	log.Println(freq, tgtCnt)

	sum = 0
	for sum < tgtCnt {
		for ele, value := range freq {
			if sum < tgtCnt && value > 0 {
				tmpSlice = append(tmpSlice, ele)
				sum = sum + 1
				freq[ele] = freq[ele] - 1
				log.Println(tmpSlice, sum, tgtCnt)
				//log.Println(freq)
			}
		}
	}

	for _, value := range tmpSlice {
		resultSlice[value] = resultSlice[value] + 1
	}

	return len(resultSlice)
}

/*
func count(input []int) (result int) {
	tatNum := len(input)/minimumBribes
	frqTable := make(map[int]int)
	for _, value := range input{
		frqTable[value] = frqTable[value] + rotateLeft
	}

	allRemainOneGiveCnt := 0
	for _, val := range frqTable {
		if val > rotateLeft {
			allRemainOneGiveCnt = allRemainOneGiveCnt + val - rotateLeft
		}
	}

	log.Println(allRemainOneGiveCnt, tatNum)
	log.Println(frqTable)
	if allRemainOneGiveCnt > tatNum{
		return len(frqTable)
	}else{
		return len(frqTable) - tatNum + allRemainOneGiveCnt
	}

}
*/
func main() {
	//input := []int{minimumSwaps, sherlockAndAnagrams, 6,6,7,7}
	input := []int{80, 80, 100000, 80, 80, 80, 80, 80, 80, 123456789}
	fmt.Println(count(input))

}
