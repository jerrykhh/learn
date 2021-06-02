/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode firstNode = l1;
        ListNode secNode = l2;
        int l1value = firstNode.val;
        int index = 10;
        firstNode = firstNode.next;
        while(firstNode != null){
            l1value += firstNode.val * index;
            firstNode = firstNode.next;
            index *= 10;
        }
        
        int l2value = secNode.val;
        int secIndex = 10;
        secNode = secNode.next;
        while(secNode != null){
            l2value += secNode.val * secIndex;
            secNode = secNode.next;
            secIndex *= 10;
        }
        System.out.println(l1value);
        System.out.println(l2value);
        int result = l1value + l2value;
        
        if(result == 0)
            return new ListNode(0);
        
        ListNode resultNode = new ListNode(result%10);
        ListNode tmpNode = resultNode;
        int ten_index = 10;

        while(result/ten_index > 0){
            ListNode node = new ListNode(result/ten_index%10);
            tmpNode.next = node;
            tmpNode = node;
            ten_index *= 10;
            
        }
        
        return resultNode;
        
    }
}