func smallerNumbersThanCurrent(nums []int) []int {
    mp := make(map[int]int)
    newNums := make( []int, len(nums))
    copy(newNums, nums)
    sort.Ints(newNums)
    for i, _ := range nums {
        if _,ok := mp[newNums[i]]; !ok{
        mp[newNums[i]] = i
        }
        
    }
    for i:=0; i < len(nums); i++ {
        newNums[i] = mp[nums[i]]
    }
    return newNums
}