/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
// @lc code=start
import "math"

func reverse(x int) int {

	// // 1.算出每一位，之后将每一位进行加

	// if x == 0 {
	// 	return 0
	// }

	// var flag int64

	// if x > 0 {
	// 	flag = 1
	// } else {
	// 	flag = -1
	// 	x = x * -1
	// }

	// vals := make([]int64, 0, 0)
	// for i := 1; x != 0; i++ {
	// 	val := int(math.Pow(10, float64(i-1)))
	// 	lval := (x % (int(math.Pow(10, float64(i))))) / val
	// 	vals = append(vals, int64(lval))
	// 	x = x - (lval * val)
	// }
	// var target int64
	// for i := 0; i < len(vals); i++ {
	// 	target = vals[i]*int64(math.Pow(10, float64(len(vals)-i-1))) + target
	// 	if target*flag < math.MinInt32 || target*flag > math.MaxInt32 {
	// 		return 0
	// 	}
	// }
	// return int(target) * int(flag)

	// 2.巧解
	// a.一个数字%10 取个位数字
	// b.一个数字/10 去掉各位数字
	var target int
	for x != 0 {

		target = target*10 + x%10
		if target < math.MinInt32 || target > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	return target
}

// @lc code=end

