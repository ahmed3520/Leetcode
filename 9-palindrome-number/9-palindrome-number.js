/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
   let val =   String(x).split('');
        let j=val.length-1;
    for(let i=0;i<val.length/2;i++){
        if(val[i]===val[j]){
            j--
            continue
        }
        return false
    }
    return true
};