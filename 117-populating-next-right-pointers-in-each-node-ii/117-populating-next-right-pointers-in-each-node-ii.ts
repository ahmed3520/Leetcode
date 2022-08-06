/**
 * Definition for Node.
 * class Node {
 *     val: number
 *     left: Node | null
 *     right: Node | null
 *     next: Node | null
 *     constructor(val?: number, left?: Node, right?: Node, next?: Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function connect(root: Node | null): Node | null {
        if (null == root) return root;
    
    let queue = [root]
    
    while (queue.length) {
        let qlen = queue.length
        let pre = null;
        for (let i = 0; i < qlen; i++) {
            let node = queue.shift()
            
            node.next = pre;
            if (null != node.right) queue.push(node.right);            
            if (null != node.left) queue.push(node.left);            
            
            pre = node;
        }
    }
    
    return root;
};