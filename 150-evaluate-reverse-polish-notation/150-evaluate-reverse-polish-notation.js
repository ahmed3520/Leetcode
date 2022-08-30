/**
 * @param {string[]} tokens
 * @return {number}
 */
var evalRPN = function(tokens) {
  let stack=[]
  for(let c of tokens){
      if(c=="+"){
          stack.push(stack.pop()+stack.pop())
      }else if(c=="-"){
          let a= stack.pop()
          let b = stack.pop()
          stack.push(b-a)
      }else if(c=="*"){
          stack.push(stack.pop()*stack.pop())
      }else if(c=="/"){
         let a = stack.pop()
         let b=  stack.pop()
         stack.push(Number(b/a| 0)) 
      }else{
          stack.push(Number(c))
      }
  }
    return stack[0]
};