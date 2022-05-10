func finalValueAfterOperations(operations []string) int {
    x := 0
    
    for i := 0; i < len(operations); i++ {
        op := operations[i]
        if op == "X++" || op == "++X" {
            x++
        } else if op == "X--" || op == "--X" {
            x--
        } else {
            //check if there any hidden test cased
        }
    }
    
    return x
}