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
 * @return {void} Do not return anything, modify root in-place instead.
 */
var recoverTree = function(root) {
    let prev = null, min = null, max = null;
    let dfs = function(root) {
        if (!root) return;
        dfs(root.left);
        if (prev != null && prev.val > root.val) {
            min = root; // potential smaller number that needs to be swapped
            if (!max) max = prev; // assured bigger number that needs to be swapped
            else return;
        }
        prev = root;
        dfs(root.right);
    }
    
    dfs(root);
    [max.val, min.val] = [min.val, max.val];
};