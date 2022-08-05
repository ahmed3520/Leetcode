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
 * @param {number} targetSum
 * @return {number[][]}
 */
var pathSum = function(root, targetSum) {
    var paths=[]
    var path=[]
    helper(root,targetSum, path, paths)
    return paths
};

 const helper = (node, total, arr,res) => {
        if(!node) return;
        if(node.val === total && !node.left && !node.right) {
            arr.push(node.val);
            res.push(arr.slice());
            arr.pop();
            return;
        }

        arr.push(node.val);

        helper(node.left, total - node.val, arr,res)
		
        helper(node.right, total - node.val, arr,res);

        arr.pop()
    }