//write a function that returns a pointer to an integer and demonstrate its usage
package main
import "fmt"

func newInt(value int) *int{
  return &value
}

func main(){
  ptr := newInt(42)
  fmt.Println("Value:", *ptr)
}
