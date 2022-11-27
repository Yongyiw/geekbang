> Problem: [66. 加一](https://leetcode.cn/problems/plus-one/description/)

[TOC]

# 思路
> 末尾+1，用carry:[1|0]表示进位，倒序操作每一位

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(n)$

# Code
```TypeScript []

function plusOne(digits: number[]): number[] {
    let carry = 1;
    const result = digits.reverse().map( d => { d += carry; carry = d >= 10 ? 1 : 0; return d % 10}).reverse();
    
    return carry ? [ 1, ...result] : result;
};

function plusOne(digits: number[]): number[] {
    let carry = 1;
    for (let i = digits.length - 1; i >=0; i--) {
        digits[i] += carry;
        carry = digits[i] >= 10 ? 1 : 0;
        digits[i] %= 10;
    }

    return carry ? [ 1, ...digits] : digits;
};
```
