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
 * @return {boolean}
 */
var isBalanced = function(root) {
    if(!root) return true
    return check(root) !=-1
};

function check(root){
    if(!root) return 0
    let left = check(root.left)
    if(left != -1){
        let right  = check(root.right)
        if(right !=-1){
            return Math.abs(left-right)<=1?Math.max(left,right)+1:-1
        }
    }
    return -1
}