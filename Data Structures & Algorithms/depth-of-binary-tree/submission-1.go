/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    m := 0
    var dfs func(node *TreeNode, height int)
    dfs = func(node *TreeNode, height int) {
        if node == nil {
            return 
        }
        m = max(m, height)
        dfs(node.Left, height+1)
        dfs(node.Right, height+1)
    }

    dfs(root, 1)
    return m
}
