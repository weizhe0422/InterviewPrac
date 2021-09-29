package main

import "fmt"

const Max = int32(^uint32(0)>>1)
const Min = int32(^Max)

func GetMin(a ...int32) int32{
	m:=Max
	for _, value := range a{
		if value < m {
			m = value
		}
	}
	return m
}

func GetMax(a ...int32) int32{
	m:=Min
	for _, value := range a{
		if value > m {
			m = value
		}
	}
	return m
}

func maxSubsetSum(arr []int32) int32 {
	var(
		dpTable []int32
		result int32
	)
	dpTable = make([]int32, len(arr))
	dpTable[0] = arr[0]
	dpTable[1] = GetMax(arr[0], arr[1])
	result = dpTable[1]

	for idx:=2; idx <= len(arr)-1; idx++{
		dpTable[idx] = GetMax(arr[idx], arr[idx]+dpTable[idx-2], dpTable[idx-2], dpTable[idx-1])
		if dpTable[idx] > result {
			result = dpTable[idx]
		}
	}
	return result
}


func main() {
	fmt.Println(maxSubsetSum([]int32{3,7,4,6,5}))
	fmt.Println(GetMax(3,7,4,6,5))
}
