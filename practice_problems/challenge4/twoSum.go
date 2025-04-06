// Leetcode 1 - Two Sum
// https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
)

func indexMap(nums []int) map[int][]int {
	result := make(map[int][]int)

	for i, num := range nums {
		result[num] = append(result[num], i)
	}

	return result
}

func twoSum(nums []int, target int) []int {
	idxMap := indexMap(nums)
	result := []int{}

	for key, value := range idxMap {
		firstIdx := value[0]
		targetNum := target - key
		targetIdx, ok := idxMap[targetNum]

		if key != targetNum && ok {
			result = []int{firstIdx, targetIdx[0]}
		} else if key == targetNum && len(value) >= 2 {
			result = []int{value[0], value[1]}
		}
	}

	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(indexMap(nums))
	fmt.Println(twoSum(nums, target))
}
