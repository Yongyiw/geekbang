> Problem: [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/)

[TOC]

# 思路
> 1 - 原地操作，不能用.filter() = shallow copy
2 - 快慢指针：快指针遍历所有N个元素，慢指针写入前自增，覆盖edge case（数组长度为0或1）
3 - size = slow + 1

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(n)$

# Code
```TypeScript []

function removeDuplicates(nums: number[]): number {
    let slow = 0;
    for (let fast = 0; fast < nums.length; fast++) {
        if (nums[slow] === nums[fast]) {
            continue;
        }

        nums[++slow] = nums[fast];
    }

    return slow + 1;
};
```
