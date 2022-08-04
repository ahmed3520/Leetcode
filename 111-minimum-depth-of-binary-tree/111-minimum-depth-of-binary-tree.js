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
var minDepth = function(root) {
    let res=0;
    if(!root) return 0
    let leftRes = minDepth(root.left)
    let rightRes = minDepth(root.right)
    console.log(leftRes, rightRes)
    return 1 + (Math.min(leftRes, rightRes) || Math.max(leftRes, rightRes))
};