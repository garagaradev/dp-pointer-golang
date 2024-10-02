//implement an interface and use a pointer to call a method defined in the interface 
package main
import "fmt"

type Shape interface {
  Area() float64
}

type Rectangle struct {
  Width, Height float64
}

func (r *Rectangle) Area() float64 {
  return r.Width * r. Height
}

func main(){
  rect := &Rectangle{Width:10, Height: 5}
  fmt.Println("Area: ", rect.Area())
}
