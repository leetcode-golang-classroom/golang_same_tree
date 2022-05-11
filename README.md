# golang_same_tree

Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```
Input: p = [1,2,3], q = [1,2,3]
Output: true

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```
Input: p = [1,2], q = [1,null,2]
Output: false

```

**Example 3:**

![https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```
Input: p = [1,2,1], q = [1,1,2]
Output: false

```

**Constraints:**

- The number of nodes in both trees is in the range `[0, 100]`.
- $`-10^4$ <= Node.val <= $10^4$`

## 解析

題目給出兩個二元樹的根結點要判斷兩個樹是否相同

兩顆樹相同的定義有兩個條件：

1. 結構相同
2. 每個相對應點的值相同

對於兩個二元樹的根結點 q, p 是相同可以歸納出以下檢驗條件

1. if q 是空值， p 必須是空值
2. if q 非空，p 必須非空， 且 p.Val = q.Val
3. q.Left 形成的子樹 與 p.Left 形成的子樹相同, q.Right 形成的子樹 與 p.Right 形成的子樹相同

## 程式碼

```go
package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	}
	if q != nil {
		if q.Val != p.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Left)
	}
	return false
}

```
## 困難點

1. 理解兩個樹是相同的條件
2. 找出遞迴關係

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity