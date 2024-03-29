/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function(root) {
    //DFS
    /*if(!root)
        return root;
        
    if(root.left != null){
        root.left.next = root.right;
        if(root.next != null)
            root.right.next = root.next.left;
    }
    
    connect(root.left);
    connect(root.right);
    return root
    */
    //BFS
    if(!root) return root
    let queue=[root]
    while(queue.length !=0){
        let next=[]
        while(queue.length != 0){
        let current = queue.shift()
        current.next = queue[0]||null
            if(current.left){
                next.push(current.left)
                next.push(current.right)
            }
        }
        queue=next
    }
return root
};