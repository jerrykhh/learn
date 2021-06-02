class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int max = 0;
        int cache = 0;
        for(int i = 0; i < nums.length; i++){
          if(nums[i] == 1){
            cache++;
          }else{
            if(cache > max){
              max = cache;
            }
            cache = 0;
          }
        }

        if(cache > max){
          max = cache;
        }

        return max;
    }
}