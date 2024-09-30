// write a function that merges two sorted linked lists into a single sorted list using pointers
package main
import "fmt"

type Node struct {
  data int 
  next *Node
}

func mergeLists(l1, l2 *Node) *Node {
  dummy := &Node{}
  current := dummy

  for l1 != nil && l2 != nil {
    if l1.data < l1.data {
      current.next = l1
      l1 = l1.next
    }else{
      current.next = l2
      l2 = l2.next
    }
    current = current.next
  }

  if l1 != nil {
    current.next = l1
  }else{
    current.next = l2
  }
  return dummy.next
}

func main(){
  l1 := &Node{1, &Node{3, &Node{5, nil}}}
  l2 := &Node{2, &Node{4, &Node{6, nil}}}

  merged:= mergeLists(l1,l2)
  for merged != nil {
    fmt.Print(merged.data, " -> ")
    merged = merged.next
  }
  fmt.Println("nil")
}
