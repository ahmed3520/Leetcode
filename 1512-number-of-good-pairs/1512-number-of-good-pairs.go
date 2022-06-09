func numIdenticalPairs(nums []int) int {
   dict := [128]int{}
    count := 0
    for _,val := range nums {
        
            count += dict[val]
            dict[val]++
        
    }
    return count
}