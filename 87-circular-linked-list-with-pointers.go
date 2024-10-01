//implement a circular linked list with pointers
package main
import "fmt"

type Node struct {
  Value int
  Next *Node
}
func main(){
  first := &Node{Value:1}
  second := &Node{Value:2}
  third := &Node{Value:3}

  first.Next = second 
  second.Next = third 
  third.Next = first

  current := first
  for i := 0; i < 6; i++ {
    fmt.Println(current.Value)
    current = current.Next
  } 

}
