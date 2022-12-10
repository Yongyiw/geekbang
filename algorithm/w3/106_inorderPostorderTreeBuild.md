> Problem: [106. 从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/)

[TOC]

# 思路
> 后序遍历root节点总在最后，可以尝试递归实现

# 解题方法
> 递归实现左右子树的构建

# 复杂度
- 时间复杂度: 
>  $O(n)$，其中 n 是树中的节点个数：

- 空间复杂度: 
> $O(n)$

# Code
```C++ []

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        // function overloads
        return buildTree(inorder, postorder, 0, inorder.size() - 1, 0, postorder.size() - 1);
    }

    /**
    * inorder[l1:r1] = inorder of sub tree (left | right), l1, r1 both inclusive
    * postorder[l2:r2] = postorder of sub tree (left | right)
    */
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder, int l1, int r1, int l2, int r2) {
        if (l1 > r1 || l2 > r2) {
            return nullptr;
        }
        
        TreeNode* root = new TreeNode(postorder[r2]);

        // Validate boundary
        if (l2 == r2 || l1 == r1) {
            return root;
        }

        // Find inorder index of the root of (sub)tree
        int rootIndex = -1;
        for (int i = l1; i <= r1; i++) {
            if (root->val == inorder[i]) {
                rootIndex = i;
                break;
            }
        }

        // inorder tree = 
        // [ l1, ... rootIndex - 1, rootIndex, rootIndex + 1 ... r1 ] 
        //   |   left sub tree   |    root   |       right sub      |

        // postorder tree = 
        // [ l2, ... size of left sub tree ... l2 + (rootIndex - l1) ... r2 - 1, root]
        //   |      left sub tree                                  | right sub | root|   

        root->left = buildTree(inorder, postorder, 
                l1,     rootIndex - 1,
                l2,     l2 + (rootIndex - l1 - 1)
        );
        root->right = buildTree(inorder, postorder, 
                rootIndex + 1,          r1, 
                l2 + (rootIndex - l1),  r2 - 1
        );

        return root;
    }
};
```
