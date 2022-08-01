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
 * @return {number}
 */
var maxDepth = function(root) {
    let res=0
    if(root==null){
        return 0
    }
   let leftRes = maxDepth(root.left)
   let rightRes = maxDepth(root.right)
    res=Math.max(leftRes,rightRes)+1
    return res
};