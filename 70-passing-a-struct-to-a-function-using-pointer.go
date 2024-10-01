//create a function that accepts a struct as a pointer and modifies one of its fields
package main
import "fmt"

type Circle struct {
  radius float64
}

func setRadius(c *Circle, r float64){
  c.radius = r
}

func main(){
  circle := &Circle{radius:5}
  fmt.Println("Before:", circle.radius)
  setRadius(circle, 10)
  fmt.Println("After:", circle.radius)
}
