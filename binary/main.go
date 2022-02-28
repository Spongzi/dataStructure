package main

import "fmt"

// BinarySearch 二分查找递归法
func BinarySearch(array []int, target int, left, right int) int {
	if left > right {
		// 及出界了，不能继续寻找了
		return -1
	}
	// 从中间开始寻找
	mid := (left + right) / 2
	middleNum := array[mid]
	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间数比目标数字大，从左边寻找
		return BinarySearch(array, target, 0, mid-1)
	} else if middleNum < target {
		// 中间数字比目标数字小，从右边寻找
		return BinarySearch(array, target, mid+1, right)
	}
	return 0
}
func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
