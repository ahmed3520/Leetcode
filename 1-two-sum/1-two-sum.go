func twoSum(nums []int, target int) []int {
    itemsMap := make(map[int]int)
    
    for i, currNum := range nums {
        if reqIndex, isPresent := itemsMap[target-currNum]; isPresent{
            fmt.Println(reqIndex, isPresent)
            return []int{reqIndex,i}
        }
        itemsMap[currNum] = i
        
    }
    return []int{}
}