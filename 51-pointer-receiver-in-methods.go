//create a struct and implement a method with a pointer receiver to modify the struct's field.
package main
import "fmt"

type Counter struct {
  count int
}

func (c *Counter) Increment(){
  c.count++
}

func main(){
  counter := Counter{count:0}
  counter.Increment()
  fmt.Println("Counter value:",counter.count)
}
