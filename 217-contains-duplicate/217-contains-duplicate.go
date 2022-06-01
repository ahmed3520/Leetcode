func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i:=0; i<len(nums); i++{
        if(i<len(nums)-1){
            if nums[i] == nums[i+1]{
                return true
            }
        }
    }
    return false
}