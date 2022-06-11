func kidsWithCandies(candies []int, extraCandies int) []bool {
       newArr := make( []bool,0) 
      newCandies := make( []int, len(candies))
      copy(newCandies, candies)
    sort.Ints(newCandies)
    for _,val := range candies {
        if val+extraCandies >= newCandies[len(newCandies)-1] {
            newArr =append(newArr, true)
        }else{
            newArr = append(newArr, false)
        }
    }
    fmt.Println(candies,newCandies)
    return newArr
}