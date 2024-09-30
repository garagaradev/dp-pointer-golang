//create a simple stack data structure using pointers
package main
import "fmt"

type Node struct {
  data int
  next *Node
}

type Stack struct {
  top *Node
}

func (s *Stack) Push(val int){
  newNode := &Node{data:val, next:s.top}
  s.top = newNode
}

func (s *Stack) Pop()int{
  if s.top == nil {
    return -1
  }

  val := s.top.data
  s.top = s.top.next
  return val
}

func main(){
  stack := Stack{}
  stack.Push(10)
  stack.Push(20)

  fmt.Println(stack.Pop())
  fmt.Println(stack.Pop())
  fmt.Println(stack.Pop())
}
