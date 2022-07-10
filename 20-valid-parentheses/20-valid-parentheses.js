/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const stack=[]
    for(let i=0; i<s.length; i++){
        let ch = s.charAt(i)
        if(ch=='('){
            stack.push(')')
        }else if(ch=='['){
            stack.push(']')
        }else if(ch=='{'){
            stack.push('}')
        }else{
            if(ch != stack.pop()){
                return false
            }
        }
    }
    return stack.length ===0
};