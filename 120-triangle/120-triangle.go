func minimumTotal(triangle [][]int) int {
  
    for r := len(triangle)-2; r>=0; r-- {
        for c := range triangle[r] {
            triangle[r][c] += min(triangle[r+1][c], triangle[r+1][c+1])
        }
    }
    return triangle[0][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
