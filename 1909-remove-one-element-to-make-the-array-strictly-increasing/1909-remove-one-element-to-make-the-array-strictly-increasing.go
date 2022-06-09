func canBeIncreasing(nums []int) bool {
    prev := nums[0]
    status := false
    for i:=1; i< len(nums); i++{
        if nums[i] <=prev{
            if(status){
                return false
            }
            status = true
            if i == 1 || nums[i] > nums[i - 2] {
                prev = nums[i]
            }
        }else{
            prev = nums[i]
        }
    }
    return true
}