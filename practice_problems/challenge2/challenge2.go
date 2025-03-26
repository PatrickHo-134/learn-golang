// Leetcode 27 - Remove element
// https://leetcode.com/problems/remove-element/description/?envType=problem-list-v2&envId=array

package main

import (
	"fmt"
)

// Returns a number of elements in nums that are not equal to val.
// Removes all occurrences of val in nums in-place
func removeElement(nums []int, val int) int {
	countElement := 0

	for _, num := range nums {
		if num != val {
			nums[countElement] = num
			countElement++
		}
	}

	return countElement
}

func main() {
	myNums := []int{3, 2, 2, 3}
	result := removeElement(myNums, 3)

	fmt.Println("Result", result)
	fmt.Println("Updated array", myNums)
}
