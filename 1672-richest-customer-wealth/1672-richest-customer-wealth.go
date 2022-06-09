func maximumWealth(accounts [][]int) int {
    max:=0
    for i:=0; i<len(accounts); i++ {
        temp:=0
        for j:=0; j < len(accounts[i]); j++{
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