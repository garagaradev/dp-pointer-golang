//implement a function that accepts a pointer to a custom type 
//and modifies its fields.
package main
import "fmt"

type Rectangle struct {
  width, height float64
}

func resize(r *Rectangle, w, h float64){
  r.width = w 
  r.height = h
}

func main(){
  rect := Rectangle{width:10, height:5}
  resize(&rect, 20, 10)
  fmt.Println("Resized Rectangle:", rect)
}
