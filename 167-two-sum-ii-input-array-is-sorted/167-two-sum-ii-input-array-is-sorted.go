func twoSum(numbers []int, target int) []int {
    i,j := 0, len(numbers)-1
    res := numbers[i] + numbers[j]
    
    for res != target {
        if res < target {
            i++
        }else  {
            j--
        }
    res = numbers[i] + numbers[j]
    }
    return []int{i+1,j+1}
}