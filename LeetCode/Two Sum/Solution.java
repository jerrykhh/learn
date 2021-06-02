import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> hash = new HashMap<>();
        int[] result = new int[2];
        for(int i = 0; i < nums.length; i++){
            int remainder = target - nums[i];
            if(hash.containsKey(remainder)){
                result[0] = hash.get(remainder);
                result[1] = i;
                break;
            }
            hash.put(nums[i], i);
        }
        return result;
    }
}
