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
        currentSum = maximum(nums[i], nums[i]+currentSum)
        maxSum = maximum(currentSum, maxSum);
    }
    return maxSum;

}