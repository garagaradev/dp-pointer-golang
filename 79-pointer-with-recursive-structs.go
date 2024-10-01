//create a recursive struct and demonstrate accessing it via pointers
package main
import "fmt"

type TreeNode struct {
  Value int
  Left *TreeNode
  Right *TreeNode
}

func printTree(node *TreeNode) {
  if node != nil {
    printTree(node.Left)
    fmt.Print(node.Value, " ")
    printTree(node.Right)
  }
}

func main(){
  root := &TreeNode{Value:10}
  root.Left = &TreeNode{Value:5}
  root.Right = &TreeNode{Value:15}

  fmt.Print("Tree values: ")
  printTree(root)
}
