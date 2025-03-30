// Leetcode 242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/description/

package main

import (
	"fmt"
	"reflect"
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChars := []rune(s)
	tChars := []rune(t)

	slices.Sort(sChars)
	slices.Sort(tChars)

	return reflect.DeepEqual(sChars, tChars)
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
