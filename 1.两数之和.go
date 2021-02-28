/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	// 1.暴力解法
	// 复杂度分析
	// 空间复杂度O(1)
	// 时间复杂度O(N^2)
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// 2.哈希表
	// 备注下：哈希表解法，第一次访问也许会错过，但是当顺序遍历到后边时，
	// 又会访问到之前的元素，所以不必担心刚开始hashtable没有而导致不能够组合。
	// 复杂度分析
	// 时间复杂度：O(N)
	// 空间复杂度：O(N)
	hashtable := make(map[int]int)
	for i, _ := range nums {
		// 这句话保证不会有重复元素
		another := target - nums[i]
		if _, ok := hashtable[another]; ok {
			return []int{i, hashtable[another]}
		}
		hashtable[nums[i]] = i
	}
	return nil
}

// @lc code=end

