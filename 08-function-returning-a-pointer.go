//write a program where a function returns a pointer to a local variable
//and explain why this works.
package main
import "fmt"

func createPointer() *int{
  x := 10
  return &x
}

func main(){
  p := createPointer()
  fmt.Println(*p)
}
