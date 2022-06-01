func maximum(first int, last int)int{
    if(first > last){
        return first
    }
    return last
}
func maxSubArray(nums []int) int {
    currentSum:=nums[0]
    maxSum:= nums[0]
    for i:=1; i< len(nums); i++ {
        if currentSum <0 {
            currentSum = nums[i]
        }else{
            currentSum += nums[i] 
        }
        fmt.Println(maxSum, currentSum)
        maxSum = maximum(currentSum,maxSum );
    }
    return maxSum;

}