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
var sumNumbers = function(root) {
    if(!root) return null
    let sum =0;
    
    function helper(root,str){
        str+=root.val
        if(!root.left&&!root.right){
        console.log("c=> ",str)
            sum+=+str
        console.log("p=> ",str)
        }
        if(root.left)
        {
            helper(root.left,str)
            console.log(str)
        }
        if(root.right) {
            helper(root.right,str)
            console.log("2=>", str)
        }
    }
    helper(root,'')
    return sum
};