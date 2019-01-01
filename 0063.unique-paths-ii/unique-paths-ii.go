package problem0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    obs := obstacleGrid
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obs[0])
    if n == 0 {
        return 0
    }

    dp := make([][]int,  m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    if obs[0][0] == 0 {
        dp[0][0] = 1
    }

    for i:=1 ; i < m ; i++{
        if obs[i][0] == 0{
            dp[i][0] = dp[i-1][0]
        }
    }

    for j := 1; j < n; j++{
        if obs[0][j] == 0{
            dp[0][j] = dp[0][j-1]
        }
    }

    for i:=1; i < m; i++ {
        for j := 1; j < n; j++{
            if obs[i][j] == 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }

    return dp[m-1][n-1]

}
