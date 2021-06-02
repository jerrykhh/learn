class Solution {
    public int[] sortedSquares(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            nums[i] = nums[i] * nums[i];
        }

        for (int i = 0; i < nums.length; i++) {
          for (int j = i + 1; j < nums.length; j++) {
            int cache = 0;
            if (nums[i] > nums[j]) {
              cache = nums[i];
              nums[i] = nums[j];
              nums[j] = cache;
            }
          }
        }

        return nums;
    }
}