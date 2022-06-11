func minOperations(nums []int, x int) int {
    sum := 0
    for _,val := range nums{
        sum+=val
    }
    target := sum - x
   start, cur_sum, max_len :=0,0,0
    f := false
    for i :=0 ; i<len(nums); i++{
        cur_sum+= nums[i]
        for start <= i && cur_sum > target {
            cur_sum -=nums[start]
            start++
        }
        if cur_sum == target{
            f = true
            max_len = max(max_len, i-start+1)
        }
    }
    if f {
        return len(nums)- max_len
    }
    return -1
    
}
func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}