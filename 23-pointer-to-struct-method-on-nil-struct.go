//demonstrate how calling a method on a nil pointer struct works without causing a panic
package main
import "fmt"

type Node struct {
  value int
}

func (n *Node) PrintValue(){
  if n == nil {
    fmt.Println("Node is nil.")
    return
  }
  fmt.Println("Node value:", n.value)
}

func main(){
  var n *Node
  n.PrintValue()
}
