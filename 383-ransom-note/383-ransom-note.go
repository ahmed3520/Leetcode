func canConstruct(ransomNote string, magazine string) bool {
    newMap := make(map[string]int)
    for i:=0;i<len(magazine); i++{
        newMap[string(magazine[i])]++
    }
    fmt.Println(newMap)
    for i:=0; i <len(ransomNote);i++{
        newMap[string(ransomNote[i])]--
         fmt.Println(newMap)
        
        if(newMap[string(ransomNote[i])] <0){
            return false
        }
    }
    return true
}