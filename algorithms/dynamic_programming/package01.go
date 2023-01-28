package dynamic_programming

// Package01 01背包
func Package01(weight, value []int, bagweight int) int {
	// dp[i][j] i表示 物品的种类，j表示 背包的可装的重量，dp[i][j]表示价值
	typeSize := len(weight)
	dp := make([][]int,typeSize)
	for i:=0;i<typeSize;i++{
		dp[i] = make([]int,bagweight+1)
	}
	// 递推公式：dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i])
	// 初始化：
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 遍历顺序
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
