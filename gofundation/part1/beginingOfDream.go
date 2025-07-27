package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// 查找唯一元素
	//unique := findUnique([]int{1, 1, 2, 2, 3, 4, 4, 5, 5})
	//fmt.Printf("unique item:%v \n", unique)

	//palindrome := isPalindrome(1221)
	//fmt.Printf("isPalindrome:%v \n", palindrome)

	//valid := isValid("{[[()]()[](([]))]}")
	//fmt.Printf("isBracket:%v\n", valid)

	/*strings := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strings)
	fmt.Printf("longest prefix:%v \n", prefix)*/

	/*digist := []int{4, 3, 2, 9}
	res := plusOne(digist)
	fmt.Printf("plusOne result:%v \n", res)*/
	/*arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(arr)
	fmt.Printf("unique nums:%v\n", n)*/

	/*intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(intervals)
	fmt.Printf("merge result:%v \n", res)*/

	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Printf("twosum res:%v \n", res)

}

/*
	136.给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。

找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func findUnique(a []int) int {
	m := make(map[int]int)
	for i := range a {
		vi := a[i]
		if _, exist := m[vi]; exist {
			delete(m, vi)
		} else {
			m[vi] = 1
		}
	}
	for k := range m {
		return k
	}
	return 0
}

/*
回文数字
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, l := 0, len(s); i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

/*
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
func isValid(s string) bool {
	// use slice
	sl := make([]string, 0)
	m := map[string]string{"[": "]", "(": ")", "{": "}"}

	for i := range s {
		u := s[i]
		//fmt.Printf("%v \n", string(u))
		switch ss := string(u); ss {
		case ")", "]", "}":
			pop := sl[len(sl)-1:]
			if m[pop[0]] != ss {
				return false
			} else {
				sl = sl[:len(sl)-1]
			}
		default:
			sl = append(sl, ss)
		}
	}
	return true

}

/*
最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/
func longestCommonPrefix(strs []string) string {
	runes := make([]byte, 0)
	for i := 0; ; i++ {
		var b byte
		for j := range strs {
			s := strs[j]
			if i > len(s)-1 {
				return string(runes)
			}
			if b == 0 {
				b = s[i]
			} else if b != s[i] {
				return string(runes)
			}
		}
		runes = append(runes, b)
	}
	return string(runes)
}

/*
基本值类型
加一
难度：简单
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
链接：https://leetcode-cn.com/problems/plus-one/
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + 1
		if d < 10 {
			digits[i] = d
			return digits
		}

		digits[i] = d % 10

		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

/*
引用类型：切片
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，
将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {
	l := len(nums)
	var i = 0
	var j = 1
	for ; j < l; j++ {
		if nums[j] == nums[j-1] {
			continue
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中
*/
func merge(intervals [][]int) [][]int {
	l := len(intervals)
	res := make([][]int, 0, l)
	if l == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var j = 1
	var last []int
	res = append(res, intervals[0])
	for ; j < l; j++ {
		last = res[len(res)-1]
		if last[1] < intervals[j][0] {
			res = append(res, intervals[j])
			continue
		} else if last[1] > intervals[j][1] {
			continue
		} else {
			last[1] = intervals[j][1]
		}
	}
	return res
}

/*
基础
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
链接：https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		left := target - nums[i]
		ix, ok := m[left]
		if ok {
			return []int{ix, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
