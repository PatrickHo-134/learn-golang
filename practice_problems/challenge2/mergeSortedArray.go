// Leetcode 88 - Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/description/

package main

import (
	"fmt"
	"sort"
)

// Merges nums1 and nums2 into a single array
// Return a single array sorted in non-decreasing order
func merge(nums1 []int, m int, nums2 []int, n int) {
	// step 1: combine all elements of 2 arrays into 1 array

	var result []int
	realNums1 := nums1[0:m]

	for _, num1 := range realNums1 {
		result = append(result, num1)
	}

	for _, num2 := range nums2 {
		result = append(result, num2)
	}

	// step 2: sort the array
	sort.Ints(result)

	// step 3: modify nums1's elements based on result's elements
	for idx, num := range result {
		nums1[idx] = num
	}

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
