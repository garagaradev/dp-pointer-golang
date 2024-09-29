// write a basic list where each node points to the next node using pointers
package main
import "fmt"

type Node struct {
  data int
  next *Node
}

type LinkedList struct {
  head *Node
}

func (list *LinkedList) Insert(val int){
  newNode := &Node{data: val}
  if list.head == nil {
    list.head = newNode
  }else {
    temp := list.head
    for temp.next != nil {
      temp = temp.next
    }
    temp.next = newNode
  }
}

func (list *LinkedList) Print(){
  temp:= list.head
  for temp != nil {
    fmt.Print(temp.data,"->")
    temp = temp.next
  }
  fmt.Println("nil")
}

func main(){
  list := LinkedList{}
  list.Insert(10)
  list.Insert(20)
  list.Insert(30)
  list.Print()
}
