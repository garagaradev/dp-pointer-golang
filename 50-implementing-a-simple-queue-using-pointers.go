//create a simple queue structure using pointers
package main
import "fmt"

type Node struct {
  data int
  next *Node
}

type Queue struct {
  front *Node
  rear *Node
}

func (q *Queue) Enqueue(val int){
  newNode := &Node{data:val}
  if q.rear == nil {
    q.front = newNode
    q.rear = newNode 
  }else{
    q.rear.next = newNode
    q.rear = newNode
  }
}

func (q *Queue) Dequeue()int{
  if q.front == nil {
    return -1
  }

  val := q.front.data
  q.front = q.front.next
  if q.front == nil {
    q.rear = nil
  }
  return val
}

func main(){
  queue := Queue{}
  queue.Enqueue(1)
  queue.Enqueue(2)
  queue.Enqueue(3)

  fmt.Println("Dequeued:",queue.Dequeue())
  fmt.Println("Dequeued:",queue.Dequeue())
  fmt.Println("Dequeued:",queue.Dequeue())
}
