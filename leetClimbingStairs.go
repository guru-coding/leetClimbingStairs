package leetClimbingStairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	var i int
	for i = 0; i <= 1 && i <= n; i++ {
		dp[i] = 1
	}
	for i = 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
