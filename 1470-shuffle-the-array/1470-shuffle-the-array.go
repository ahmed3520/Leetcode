func shuffle(nums []int, n int) []int {
    x := nums[:n]
    y := nums[n:]
    var newArray []int
    for i :=0; i <len(x); i++{
        newArray = append(newArray, x[i])
        newArray = append(newArray, y[i])
    }
    return newArray
}