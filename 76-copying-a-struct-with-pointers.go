//write a function that copies a struct using pointers
package main
import "fmt"

type Point struct {
  x, y int
}

func copyPoint(original *Point) Point{
  return *original
}

func main(){
  p1 := &Point{x:5, y:10}
  p2 := copyPoint(p1)
  fmt.Println("Original:",*p1)
  fmt.Println("Copy:", p2)
}
