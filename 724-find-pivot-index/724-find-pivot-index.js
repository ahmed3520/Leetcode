/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function(nums) {
    let arrSum = 0;
    let sum=0;
    for(let i=0; i<nums.length; i++){
        arrSum+=nums[i]
    }
    console.log(arrSum)
    
    for(let i=0;i<nums.length;i++){
        if(2 * sum + nums[i] == arrSum)
                return i;
            sum += nums[i];
   
        
    }
    return -1;
};