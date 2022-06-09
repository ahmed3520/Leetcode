func maximumWealth(accounts [][]int) int {
    max:=0
    for i := range accounts {
        temp:=0
        for j := range accounts[i]{
         temp+= accounts[i][j]   
        }
        max = getMaximum(max,temp)
    }
    return max
}
func getMaximum(a,b int)int{
    if a >b{
        return a
    }
    return b
}