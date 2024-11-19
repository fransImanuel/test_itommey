package main

import "fmt"

func main() {
	fmt.Println(SumArr([]int{2, 7, 11, 15}, 9))
	fmt.Println(SumArr([]int{1, 2, 3, 4}, 8))
}

func SumArr(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return true
			}
		}
	}

	return false
}
