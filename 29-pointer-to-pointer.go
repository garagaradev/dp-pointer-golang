// Write a program that uses a pointer to a pointer (double pointer) to modify an integer.
package main
import "fmt"

func modifyDoublePointer(pp **int){
  **pp = 20
}

func main(){
  x := 10
  p := &x
  modifyDoublePointer(&p)
  fmt.Println(x)
}
