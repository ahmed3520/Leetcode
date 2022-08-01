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
var isSymmetric = function(root) {
    if(root==null) return true
    return isMiror(root.left, root.right)
};

function isMiror(l,r){
    if(r==null && l == null) return true
    if(r==null || l == null) return false
    return (l.val==r.val)&&isMiror(l.left,r.right) && isMiror(l.right, r.left)
}