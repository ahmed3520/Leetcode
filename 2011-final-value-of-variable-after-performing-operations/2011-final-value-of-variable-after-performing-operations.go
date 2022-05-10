func finalValueAfterOperations(operations []string) int {
    x:=0
    //improved solution 
    for i:=0; i< len(operations); i++ {
        if (operations[i] == "X++" || operations[i] == "++X"){
            x++
        }else {
            x--
        }
}
        return x
}