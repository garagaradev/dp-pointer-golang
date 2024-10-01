//implement a simple linked list and create a function to add new node 
//at the beginning using pointers.
package main
import "fmt"

type Node struct {
  value int
  next *Node
}

func addNode(head **Node, value int){
  newNode := &Node{value: value, next: *head }
  *head = newNode
}

func printList(head *Node){
  current := head
  for current != nil {
    fmt.Println(current.value, " ")
    current = current.next
  }
  fmt.Println()
}

func main(){
  var head *Node
  addNode(&head, 10)
  addNode(&head, 20)
  addNode(&head, 30)
  printList(head)
}
