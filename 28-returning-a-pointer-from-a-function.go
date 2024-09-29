//write a program where a function returns a pointer to a local variable
package main
import "fmt"

func createPointer() *int{
  x := 100
  return &x
}

func main(){
  p := createPointer()
  fmt.Println(*p)
}
