package tanxin

import (
	"math"
	"sort"
)

/*
	假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

	对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。

	如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

	输入: g = [1,2,3], s = [1,1]
	输出: 1
	解释:
	你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
	虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
	所以你应该输出1。

*/
// FindContentChildren 455. 分发饼干
func FindContentChildren(g []int, s []int) int {
	if len(s) == 0{
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	i,j := len(g) - 1,len(s) - 1
	res := 0
	for j>=0 && i>=0{
		if s[j] >= g[i]{
			res++
			j--
			i--
			continue
		}
		i--
	}

	return res
}

/*
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。

例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。

给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。

示例 1:

    输入: [1,7,4,9,2,5]
    输出: 6
    解释: 整个序列均为摆动序列。

*/

// wiggleMaxLength 376. 摆动序列
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 2{
		return len(nums)
	}

	result,preGap,curGap := 1,0,0
	for i:=0;i<len(nums) - 1;i++{
		curGap = nums[i+1] - nums[i]
		if (preGap >= 0 && curGap < 0) || (preGap <= 0 && curGap > 0){
			result++
			preGap = curGap
		}
	}

	return result
}

// maxSubArray 53. 最大子序和
func maxSubArray(nums []int) int {
	result := math.MinInt64
	count := 0

	for i:=0;i<len(nums);i++{
		count += nums[i]

		if result < count{
			result = count
		}

		if count < 0{
			count = 0
		}
	}

	return result
}

/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

    输入: [7,1,5,3,6,4]
    输出: 7
    解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4。随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

*/

// maxProfit 122.买卖股票的最佳时机II
func maxProfit(prices []int) int {
	temp := make([]int,0,len(prices) - 1) // 存放每天的利润，除了第一天之外
	for i:=1;i<len(prices);i++{
		temp = append(temp,prices[i] - prices[i - 1])
	}

	res := 0
	for _,v := range temp{
		if v > 0{
			res += v
		}
	}

	return res
}

// canJump 55. 跳跃游戏
func canJump(nums []int) bool {
	maxScale := 0
	end := len(nums) - 1

	if end < 1{
		return true
	}

	for i:=0;i<=maxScale;i++{
		curMaxScale := nums[i] + i // 当前能达到的最大范围
		if maxScale < curMaxScale{
			maxScale = curMaxScale
		}
		if maxScale >= end{
			return true
		}
	}


	return false
}

// jump 45.跳跃游戏II
func jump(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 1{
		return 0
	}

	curScale,nextScale := 0,0
	step := 0
	for i:=0;i<numsLen;i++{
		temp := nums[i] + i
		if temp > nextScale{
			nextScale = temp
		}

		if i == curScale{
			if curScale != numsLen - 1{
				step++
				curScale = nextScale
				if curScale >= numsLen - 1{
					return step
				}
			}else{
				return step
			}
		}
	}

	return step
}

/*
给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）

以这种方式修改数组后，返回数组可能的最大和。

示例 1：

    输入：A = [4,2,3], K = 1
    输出：5
    解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。

*/
// largestSumAfterKNegations 1005.K次取反后最大化的数组和
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	lastIndex := len(nums) - 1
	for i:=0;i<=lastIndex;i++{
		if nums[i] < 0 && k > 0{
			nums[i] = -nums[i]
			k--
		}
	}


	if k%2 == 1 {
		nums[lastIndex] = -nums[lastIndex]
	}


	res := 0
	for i:=0;i<=lastIndex;i++{
		res += nums[i]
	}

	return res
}

/*
	在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明:

    如果题目有解，该答案即为唯一答案。
    输入数组均为非空数组，且长度相同。
    输入数组中的元素均为非负数。

	示例 1: 输入:

    gas = [1,2,3,4,5]
    cost = [3,4,5,1,2]

	输出: 3 解释:

    从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
    开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
    开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
    开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
    开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
    开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
    因此，3 可为起始索引。
*/

// canCompleteCircuit 134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	totalSum,curSum := 0,0

	nums := len(gas)
	start := 0
	for i:=0;i<nums;i++{
		totalSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]

		if curSum < 0{
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0{
		return -1
	}

	return start
}

/*
师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

    每个孩子至少分配到 1 个糖果。
    相邻的孩子中，评分高的孩子必须获得更多的糖果。

那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

    输入: [1,0,2]
    输出: 5
    解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。

*/

// candy 分发糖果 135.
func candy(ratings []int) int {
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	need := make([]int, len(ratings))
	sum := 0
	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 0; i < len(ratings) - 1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings)-1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = findMax(need[i-1], need[i]+1)
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}
func findMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// lemonadeChange 860.柠檬水找零
func lemonadeChange(bills []int) bool {
	fiveNums,tenNums := 0,0

	for _,v := range bills{
		if v == 5{
			fiveNums++
		}else if v == 10{
			tenNums++
			fiveNums--
		}else{
			if tenNums <= 0{
				fiveNums -= 3
			}else{
				tenNums--
				fiveNums--
			}
		}

		if fiveNums < 0 || tenNums < 0{
			return false
		}
	}

	return true
}

/*
	给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意: 可以认为区间的终点总是大于它的起点。 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

示例 1:

    输入: [ [1,2], [2,3], [3,4], [1,3] ]
    输出: 1
    解释: 移除 [1,3] 后，剩下的区间没有重叠。

*/
// eraseOverlapIntervals 435. 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals,func(i,j int)bool{
		return intervals[i][1] < intervals[j][1]
	})

	res := 1
	for i:=1;i<len(intervals);i++{
		if intervals[i][0] >= intervals[i - 1][1]{
			res++
		}else{
			intervals[i][1] = min(intervals[i][1],intervals[i-1][1])
		}
	}

	return len(intervals) - res
}

func min(a,b int)int{
	if a < b{
		return a
	}

	return b
}