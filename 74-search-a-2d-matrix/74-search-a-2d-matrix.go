func searchMatrix(matrix [][]int, target int) bool {
 
    row,col:=len(matrix),len(matrix[0])
    top:=0
    bot:=row-1
    
    for top<=bot{
        row=(top+bot)/2
        if target > matrix[row][col-1]{
            top=row+1
        }else if target<matrix[row][0]{
            bot=row-1
        }else{
            break
        }
    }
    
    if ! (top <=bot){
        return  false
    }
    
    row=(top+bot)/2
    l,r:=0,col-1
    for l<=r{
        m:=(l+r)/2
        if target> matrix[row][m]{
            l=m+1
        } else if target< matrix[row][m]{
            r=m-1
        } else{
            return true
        }
        
    }
    return false
    
}