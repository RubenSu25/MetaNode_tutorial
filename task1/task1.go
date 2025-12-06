package task1

import (
	"fmt"
	"slices"
)

// 只出现一次的数字
func SingleNumber(nums []int) int {
	count_map := make(map[int]int)
	for _, v := range nums {
		if count, exist := count_map[v]; exist {
			count_map[v] = count + 1
		} else {
			count_map[v] = 1
		}
	}
	for key, value := range count_map {
		if value == 1 {
			return key
		}
	}
	return 0
}

// 回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := fmt.Sprintf("%d", x)
	r := []rune(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true

}

// 有效的括号
func IsValid(s string) bool {
	// 如果长度为奇数，直接返回 false
	if len(s)%2 != 0 {
		return false
	}

	// 定义右括号 -> 左括号的映射
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 初始化一个栈
	stack := []rune{}

	// 遍历字符串中的每个字符
	for _, ch := range s {
		// 如果是左括号，压入栈中
		if mp[ch] == 0 { // 说明不是右括号
			stack = append(stack, ch)
		} else { // 是右括号
			// 栈为空 或者 栈顶元素不匹配当前右括号，返回 false
			if len(stack) == 0 || stack[len(stack)-1] != mp[ch] {
				return false
			}
			// 匹配成功，弹出栈顶
			stack = stack[:len(stack)-1]
		}
	}

	// 最终栈应为空，表示所有括号都匹配
	return len(stack) == 0

}

// 最长前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	s0 := strs[0]
	for j, c := range s0 { // 从左到右
		for _, s := range strs { // 从上到下
			if j == len(s) || s[j] != byte(c) { // 这一列有字母缺失或者不同
				return s0[:j] // 0 到 j-1 是公共前缀
			}
		}
	}
	return s0
}

// 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i > 0; i = i - 1 {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
		}
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// 去重
func RemoveDuplicates(nums []int) int {
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[result] = nums[i]
			result++
		}
	}
	return result
}

// 合并重叠区间
func Merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return

}

// 两数之和
func TwoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
