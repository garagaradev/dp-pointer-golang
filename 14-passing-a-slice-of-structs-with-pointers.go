// Write a program that passes a slice of structs with pointers
// and modifies the slice elements through pointers
package main
import "fmt"

type Point struct {
  x *int
}

func modifyPoints (points []Point) {
  for _, p := range points {
    *p.x += 10
  }
}

func main(){
  x1, x2 := 5, 15
  points := []Point{{&x1},{&x2}}

  fmt.Println("Before:", *points[0].x, *points[1].x)
  modifyPoints(points)
  fmt.Println("After:", *points[0].x, *points[1].x)
}
