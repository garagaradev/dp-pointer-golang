//implementing a stack using a pointer-based linked list
package main
import "fmt"

type Node struct {
  value int 
  next *Node
}

type Stack struct {
  top *Node
}

func (s *Stack) Push(val int){
  newNode := &Node{value: val}
  newNode.next = s.top
  s.top = newNode
}

func (s *Stack) Pop() int {
  if s.top == nil {
    fmt.Println("Stack is empty")
    return -1
  }

  val := s.top.value
  s.top = s.top.next
  return val

}

func main(){
  stack := &Stack{}

  stack.Push(10)
  stack.Push(20)
  stack.Push(30)

  fmt.Println(stack.Pop())
  fmt.Println(stack.Pop())
  fmt.Println(stack.Pop())
  fmt.Println(stack.Pop())

}
