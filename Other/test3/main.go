package main

import (
	"fmt"
	"log"
)

func Solution(A []int) int {
	// write your code in Go rotateLeft.sherlockAndAnagrams
	var (
		price                        []int
		idx, j, oneDTick, sevenDTick int
	)

	price = make([]int, len(A))
	price[0] = 2
	sevenDTick = 7

	for idx = 1; idx < len(A); idx++ {
		log.Println(price, idx)
		oneDTick = 2 + price[idx-1]
		j = idx - 1

		for j >= 0 && A[j] >= A[idx]-6 {

			j--
			log.Println("--", A[j], A[idx], j)
		}

		if j >= 0 {
			sevenDTick += price[j]
		}

		if oneDTick > sevenDTick {
			price[idx] = sevenDTick
		} else {
			price[idx] = oneDTick
		}
	}

	log.Println("final", price)
	return price[len(price)-1]

}
func main() {

	fmt.Println(Solution([]int{1, 2, 4, 5, 7, 29, 30}))
}
