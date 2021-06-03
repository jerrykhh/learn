class Solution {
    public int reverse(int x) {
        String s = Integer.toString(x);
        String substring = "";
        if(x >= 0)
            for(int i = s.length() -1 ; i > -1; i--)
                substring += s.charAt(i);
        else
             for(int i = s.length() -1 ; i > 0; i--)
                substring += s.charAt(i);
        
        
        return (x > 0) ? Integer.parseInt(substring):-Integer.parseInt(substring);
    }
}