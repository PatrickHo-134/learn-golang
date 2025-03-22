// Leetcode-26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=problem-list-v2&envId=array

package main

import (
	"fmt"
)

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}

	return false
}

func removeDuplicates(nums []int) int {
	var result []int
	uniqueCount := 0

	for _, num := range nums {
		if !contains(result, num) {
			result = append(result, num)
			nums[uniqueCount] = num
			uniqueCount++
		}

	}

	return uniqueCount
}

func main() {
	myNums := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}

	fmt.Println(removeDuplicates(myNums))
	fmt.Println(myNums)
}
