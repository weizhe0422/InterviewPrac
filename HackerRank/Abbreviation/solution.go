package main

import (
	"fmt"
)


func abbreviation(a string, b string) string {
	// Write your code here
	rowLen := len(a) + 1
	colLen := len(b) + 1
	isValid := make([][]bool, rowLen)
	for i := 0; i < len(isValid); i++ {
		isValid[i] = make([]bool, colLen)
	}

	isValid[0][0] = true
	j := 1
	for ;j < colLen; j++ {
		isValid[0][j] = false
	}

	hasUpperCaseString := false
	i := 1
	for ; i < rowLen; i++ {
		idxA := i - 1
		if a[idxA] >= 65 && a[idxA] <= 90 || hasUpperCaseString {
			hasUpperCaseString = true
			isValid[i][0] = false
		} else {
			isValid[i][0] = true
		}
	}

	for i = 1; i < rowLen; i++ {
		for j = 1; j < colLen; j++ {
			idxA := i - 1
			idxB := j - 1

			if a[idxA] == b[idxB] {
				isValid[i][j] = isValid[i-1][j-1]
			} else if a[idxA]-32 == b[idxB] {
				isValid[i][j] = isValid[i-1][j-1] || isValid[i-1][j]
			} else if a[idxA] >= 65 && a[idxA] <= 90 {
				isValid[i][j] = false
			} else {
				isValid[i][j] = isValid[i-1][j]
			}
		}
	}

	if isValid[rowLen-1][colLen-1] {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(abbreviation("RDWPJPAMKGRIWAPBZSYWALDBLDOFLWIQPMPLEMCJXKAENTLVYMSJNRJAQQPWAGVcGOHEWQYZDJRAXZOYDMNZJVUSJGKKKSYNCSFWKVNHOGVYULALKEBUNZHERDDOFCYWBUCJGbvqlddfazmmohcewjg","RDPJPAMKGRIWAPBZSYWALDBLOFWIQPMPLEMCJXKAENTLVYMJNRJAQQPWAGVGOHEWQYZDJRAXZOYDMNZJVUSJGKKKSYNCSFWKVNHOGVYULALKEBUNZHERDOFCYWBUCJG"))
	//fmt.Println(abbreviation("beFgH","EFG"))
	a:="az"
	fmt.Println(a[0],a[1])
}
