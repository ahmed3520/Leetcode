/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function(s) {
    helper(  s, 0,  s.length - 1)
};

const helper = (s,left, right) => {
        if (left >= right) return
        [s[left], s[right]] = [s[right], s[left]]
        helper( s, left+1, right-1)
}