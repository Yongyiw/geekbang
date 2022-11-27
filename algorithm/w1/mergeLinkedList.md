> Problem: [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/description/)

[TOC]

# 思路
> 链表归并

# 解题方法
> 用虚拟节点start记录初始节点
current表示当前结果链表的生成位置
归并完必定有一个链表指针为null，另外一个非空链表要连到结果链表末尾

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(n)$

# Code
```TypeScript []

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    let head = new ListNode(-1, null);
    let current = head;
    
    // Merging
    while(list1 !== null && list2 !== null) {
        if (list1.val > list2.val) {
            current.next = list2;
            list2 = list2.next;
        } else {
            current.next = list1;
            list1 = list1.next;
        }

        current = current.next;
    }

    // Dealing with remains
    current.next = list1 ? list1 : list2;   

    return head.next;
};
```
