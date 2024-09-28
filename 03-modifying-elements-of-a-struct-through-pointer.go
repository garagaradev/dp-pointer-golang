//write a program that modifies the fields of a struct using a pointer to the struct
package main
import "fmt"

type Rectangle struct {
  length int
  width int
}

func modifyRectangle(r *Rectangle){
  r.length = 15
  r.width = 25
}

func main(){
  rect := Rectangle{10,20}
  fmt.Println("Before:",rect)
  modifyRectangle(&rect)
  fmt.Println("After:",rect)
}
