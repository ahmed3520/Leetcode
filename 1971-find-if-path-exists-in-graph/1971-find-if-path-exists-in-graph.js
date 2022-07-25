/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number} source
 * @param {number} destination
 * @return {boolean}
 */
var validPath = function(n, edges, source, destination) {
   let adj = new Array(n).fill().map(_=>new Set());
    for(const [from, to] of edges) {
        adj[from].add(to);
        adj[to].add(from);
    }
    let seen = new Set();
    console.log(seen, adj)
    return dfs(adj,source,destination,seen)
};

function dfs(adj,source,destination,seen){
    if(source == destination){
        return true
    }
    for(const neighbour of adj[source]){
        console.log(neighbour)
        if(seen.has(neighbour))
            continue
        seen.add(neighbour)
        if(dfs(adj, neighbour, destination,seen)){
            return true
        }
    }
    return false
}