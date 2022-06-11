//2nd sol
func kidsWithCandies(candies []int, extraCandies int) []bool {
       newArr := make( []bool,0) 
    maximumNumber :=getTheMaxNumberInCandies(candies) 
    for _,val := range candies {
        if val+extraCandies >= maximumNumber{
            newArr =append(newArr, true)
        }else{
            newArr = append(newArr, false)
        }
    }
    return newArr
}
func getTheMaxNumberInCandies(nums []int) int {
    temp := 0
    for _, val := range nums {
        if val > temp {
            temp = val
        }
    }
    return temp
}