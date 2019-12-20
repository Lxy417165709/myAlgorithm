package main

/*
	给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。
*/
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			return nums[l] // l == r 时表示找到了
		}
		mid := (l + r) / 2

		// 这只是为了简化一下，让接下来只判断 mid 左右两边是偶数长度的情况。
		if mid%2 == 1 {
			mid --
		}
		if nums[mid+1] == nums[mid] {
			l = mid + 2
		} else {
			r = mid
		}
	}
	return -1
}

/*
	题目链接:
		https://leetcode-cn.com/problems/single-element-in-a-sorted-array/			有序数组中的单一元素
*/
/*
	总结
	1. 这题之前也遇到过类似，就是一个数组中，除1个数出现一次外，其他的都出现2次，求这个唯一的数。
		当时采用了异或AC。但是这题加了一个条件，就是已经排序好的，而且时间复杂度要求为O(logn),
		那么我们可以根据这个唯一的数所位于的数组长度进行二分(该数组如果有唯一的数的话，那么它的长度为偶数
		)，根据这个性质，我们就可以解决了。
	2. 以上是简化过的代码，之前写的有点复杂，但是思路比较清晰。
*/
