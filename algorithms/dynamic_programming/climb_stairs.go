package dynamic_programming

func ClimbStairs(n int) int {
	// dp 是什么：dp[i] 爬到第 i 层就有 dp[i] 种方法
	// 递推公式：dp[i] = dp[i-1] + dp[i-2]
	// 如何初始化 dp[1] = 1 dp[2] = 2 dp[3]= dp[1] + dp[2]
	// 举例子推导
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
