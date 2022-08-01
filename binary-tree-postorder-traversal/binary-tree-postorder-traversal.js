/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var postorderTraversal = function(root) {
     if(root==null)return []
    const result=[]
    const stack=[root]
    while(stack.length >0){
        const current = stack[stack.length-1]
         if(current.left){
              stack.push(current.left)
            current.left=null
        }
       else if(current.right){
            stack.push(current.right)
            current.right=null
        }else{
            result.push(stack.pop().val)
        }
    }
    return result
};