func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

func maximumWealth(accounts [][]int) int {
    wealth := 0
    
    for i := range accounts {
        sum := 0
        for j := range accounts[i] {
            sum += accounts[i][j]
        }
        wealth = max(wealth, sum)
    }
    
    return wealth
}