func numJewelsInStones(jewels string, stones string) int {
    fmt.Println("a"=="A");
        dict := [128]bool{}
    for _,j := range jewels {
        dict[j] = true
    }
    res := 0
    for _, stone := range stones {
        if dict[stone]{
            res++
        }
    }
    return res
}