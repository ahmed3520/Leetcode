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
var inorderTraversal = function(root) {
    if(root==null)return []
    const result=[]
    const stack=[root]
    while(stack.length >0){
        const current = stack[stack.length-1]
        if(current.left){
            stack.push(current.left)
            current.left=null
        }else{
            result.push(stack.pop().val)
            if(current.right) stack.push(current.right)
        }
    }
    return result
};