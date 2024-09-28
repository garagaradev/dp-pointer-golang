//write a program that demonstrates the use of pointers with method in structs
package main
import "fmt"

type Counter struct {
  value int
}

func (c *Counter) Increment(){
  c.value++
}

func main(){
  c := &Counter{value:3}
  fmt.Println("Before:",c.value)
  c.Increment()
  fmt.Println("After:",c.value)
}
