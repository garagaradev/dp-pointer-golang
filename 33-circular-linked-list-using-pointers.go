//create a circular linked list where the last node points to the first node.
package main
import "fmt"

type Node struct {
  data int
  next *Node
}

type CircularList struct {
  head *Node
}

func (list *CircularList) Insert(val int){
  newNode := &Node{data: val}
  if list.head == nil {
    list.head = newNode
    list.head.next = list.head
  }else {
    temp := list.head
    for temp.next != list.head {
      temp = temp.next
    }
    temp.next = newNode
    newNode.next = list.head
  }
}

func (list *CircularList) Print(){
  if list.head == nil {
    return
  }

  temp := list.head
  for {
    fmt.Print(temp.data, " -> ")
    temp = temp.next
    if temp == list.head {
      break
    }
  }
  fmt.Println("(back to head)")
}

func main(){
  list := CircularList{}
  list.Insert(1)
  list.Insert(2)
  list.Insert(3)
  list.Print()
}
