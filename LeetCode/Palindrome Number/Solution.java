class Solution {
    public boolean isPalindrome(int x) {
        int value = x;
        if(x < 0)
            return false;
        int palindrome = 0;
        while(x != 0){
            int tmp = x % 10;
            palindrome = palindrome * 10 + tmp;
            x /= 10;
        }
        return (palindrome == value)? true: false;
    }
}