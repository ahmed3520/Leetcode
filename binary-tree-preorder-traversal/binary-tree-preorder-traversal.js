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
var preorderTraversal = function(root) {
    if(root === null) return []
    const result=[]
    const stack=[root]
    while(stack.length > 0){
        const current  = stack.pop()
        result.push(current.val)
        console.log('current val=> ', current.val)
        console.log('current left=> ', current.left)
        console.log('current right=> ', current.right)
        if(current.right)stack.push(current.right)
                if(current.left)stack.push(current.left)

    }
    console.log(result)
    return result
};