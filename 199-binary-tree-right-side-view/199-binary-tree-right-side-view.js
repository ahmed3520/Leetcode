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
var rightSideView = function(root) {
    let data=[]
    function helper(root,result,depth){
        if(!root) return
        if(depth==result.length){
            result.push(root.val)
        }
        helper(root.right, result,depth+1)
        helper(root.left, result,depth+1)
    }
    helper(root,data,0)
    return data
};