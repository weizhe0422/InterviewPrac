package main

import (
	"fmt"
	"log"
	"regexp"
)

func solution(T []string, R []string) int {
	var (
		frqTable    map[string][]int
		testGrpPos  int
		correct     bool
		resultScore int
		perScore    int
		correctCnt  int
	)

	r := regexp.MustCompile(`[0-9]+`)

	frqTable = make(map[string][]int, len(T))
	for idx, value := range T {
		testGrpPos = r.FindStringIndex(value)[1]
		frqTable[value[0:testGrpPos]] = append(frqTable[value[0:testGrpPos]], idx)
	}
	log.Println(frqTable)

	perScore = 100 / len(frqTable)
	log.Println("perScore:", perScore)

	for testName, test := range frqTable {
		correct = true
		for _, testNo := range test {
			if R[testNo] == "fail" {
				correct = false
				break
			}
		}

		if correct {
			log.Println("test", testName)
			correctCnt++
		}
	}

	resultScore = correctCnt * perScore

	return resultScore

}

func main() {

	//test grp1
	//T:=[]string{"test1a","test2","test1b","test1c","test3"}
	//R:=[]string{"fail","ok","fail","ok","fail"}

	//test grp2
	T := []string{"codility1", "codility3", "codility2", "codility4b", "codility4a"}
	R := []string{"fail", "ok", "ok", "fail", "ok"}

	fmt.Println(solution(T, R))

}
