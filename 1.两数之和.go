/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	sli := []int{}
	for i := 0; i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+ nums[j] == target{
				sli = append(sli,i)
				sli = append(sli,j)
			}
		}
	}
	return sli
}
// @lc code=end

