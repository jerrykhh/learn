public class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode solution = new ListNode();
        ListNode tmpNode = solution;
        int carry = 0;
        while (l1 != null || l2 != null) {
            int value = 0;
            if (l1 != null && l2 != null) {
                value = l1.val + l2.val + carry;
                l1 = l1.next;
                l2 = l2.next;

            } else if (l1 != null) {
                value = l1.val + carry;
                l1 = l1.next;
            } else if (l2 != null) {
                value = l2.val + carry;
                l2 = l2.next;
            }
            carry = 0;
            
            if (value > 9) {
                carry = value / 10;
                value %= 10;
            }

            tmpNode.next = new ListNode(value);
            tmpNode = tmpNode.next;
        }
        
        if(carry > 0)
            tmpNode.next = new ListNode(carry);

        return solution.next;

    }
}
