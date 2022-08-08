/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
    let lis= [nums[0]]
    for(let i=1; i<nums.length;i++){
        let current = nums[i]
        let prev = lis[i-1]
        if(current>prev){
            lis.push(current)
        }else{
            let j=0;
            while(current>lis[j]){
                j++
            }
            lis[j]=current
        }
    }
    return lis.length
};