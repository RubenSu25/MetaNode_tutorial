package task1

import (
	"fmt"
	"testing"
)

func TestTask1(t *testing.T) {
	nums := []int{1, 2, 2}
	fmt.Printf("singleNumber(nums): %v\n", SingleNumber(nums))

	fmt.Printf("isPalindrome(): %v\n", IsPalindrome(23432))

	fmt.Printf("isValid(): %v\n", IsValid("(){}"))

	strs := []string{"carc", "carcecar", "car"}
	fmt.Printf("longestCommonPrefix(): %v\n", LongestCommonPrefix(strs))

	digits := []int{9, 9}
	fmt.Printf("plusOne: %v\n", PlusOne(digits))

	slice1 := []int{2, 3, 4, 4, 6}
	fmt.Printf("RemoveDuplicates: %v\n", RemoveDuplicates(slice1))

	slice2 := [][]int{{1, 2}, {3, 5}, {4, 7}}
	fmt.Printf("Merge: %v\n", Merge(slice2))

	slice := []int{2, 3, 4, 6}
	fmt.Printf("twoSum(): %v\n", TwoSum(slice, 7))

}
