/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let maxInteger = Math.pow(2,31)-1;
    var digit, result = 0

    while( x ){
        digit = x % 10  //  Get right-most digit. Ex. 123/10 → 12.3 → 3
        result = (result * 10) + digit  //  Ex. 123 → 1230 + 4 → 1234
        x = x/10|0  //  Remove right-most digit. Ex. 123 → 12.3 → 12
    }  
  if(result>maxInteger || result< -maxInteger){
      return 0
  }
    return result
};