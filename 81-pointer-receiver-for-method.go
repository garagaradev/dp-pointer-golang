//create a method with a pointer receiver to modify the struct's fields
package main
import "fmt"

type Circle struct {
  Radius float64
}

func (c *Circle) SetRadius(r float64) {
  c.Radius = r
}

func main(){
  circle := &Circle{Radius:5}
  fmt.Println("Before:", circle.Radius)
  circle.SetRadius(10)
  fmt.Println("After:", circle.Radius)
}
