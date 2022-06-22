/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
    let m=[]
    let m2=[]
    
    
    for(let i=0; i<s.length; i++){
        if(m[s[i]] != m2[t[i]]){

            return false
        }
        m[s[i]] = i + 1;
        m2[t[i]] = i + 1;
        
    }
    return true
};