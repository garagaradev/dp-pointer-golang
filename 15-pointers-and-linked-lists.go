//write a program to implement a simple singly linked list with pointers.
package main
import "fmt"

type Node struct {
  value int
  next *Node
}

func printList(head *Node){
  for head != nil {
    fmt.Print(head.value)
    head = head.next
    fmt.Println()
  }
}

func main(){
  head := &Node{value:1}
  second := &Node{value:2}
  third := &Node{value:3}
  fmt.Println("mem addr:",head)
  fmt.Println("mem addr:",second)
  fmt.Println("mem addr:",third)
  
  head.next = second
  second.next = third

  fmt.Println(head)
  fmt.Println(second)
  fmt.Println(third)

  printList(head)

}
