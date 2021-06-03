class Solution {
    public int reverse(int x) {
        int solution = 0;

        while (x != 0) {
            int tmp = x % 10;

            if (solution > Integer.MAX_VALUE/10 || solution < Integer.MIN_VALUE/10) {
                return 0;
            }
            solution = solution * 10 + tmp;
            x /= 10;
        }
       
        return solution;
    }
}