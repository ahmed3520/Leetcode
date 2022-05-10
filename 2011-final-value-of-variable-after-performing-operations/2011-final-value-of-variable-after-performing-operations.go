func finalValueAfterOperations(operations []string) int {
    x:=0
    mark:=0
    for i:=0; i< len(operations); i++ {
        temp := operations[i]
        if temp[0] != '-' || temp[0] != '+' {
            mark = 1 
        }
        if temp[mark]== '-' {
            fmt.Println(temp)
            x--
            fmt.Println(x)
            
        }else {
        x++
        }  
    }
    return x
}