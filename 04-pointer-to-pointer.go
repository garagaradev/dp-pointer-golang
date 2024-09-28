//write a program that demonstrates the use of a pointer to a pointer (double pointer)
package main
import "fmt"

func modifyDoublePointer(p **int){
  **p = 100
}

func main(){
  x := 10
  p := &x
  pp := &p
  fmt.Println("Before:",x)
  modifyDoublePointer(pp)
  fmt.Println("After:",x)
}
