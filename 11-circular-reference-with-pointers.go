//create a circular reference between two structs using pointers
package main
import "fmt"

type Node struct {
  value int
  next *Node
}

func main(){
  node1 := &Node{value: 1}
  node2 := &Node{value: 2}

  node1.next = node2
  node2.next = node1

  fmt.Println("node1:",node1.value,"node2:",node1.next.value)
  fmt.Println("node2:",node2.value,"node1:",node2.next.value)
}
