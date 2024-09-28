//write a program to demonstrate assigning a pointer to an interface type
//and accessing its underlying value.
package main
import "fmt"

type Shape interface {
  Area() int
}

type Square struct {
  side int
}

func (s *Square) Area() int {
  return s.side * s.side
}

func main(){
  sq := Square{side:4}
  var s Shape = &sq
  fmt.Println(s.Area())
}
