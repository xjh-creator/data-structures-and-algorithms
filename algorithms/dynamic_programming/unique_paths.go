package dynamic_programming

import "fmt"

func UniquePaths(m,n int) int {
	// dp[i][j] 含义：表示从 （0,0）出发到 （i,j）有 dp[i][j] 不同的路径
	// 递推公式：dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	// for (int i = 0; i < m; i++) dp[i][0] = 1;
	// for (int j = 0; j < n; j++) dp[0][j] = 1;
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
		dp[i][0] = 1
	}
	for i:=0;i<n;i++{
		dp[0][i] = 1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
