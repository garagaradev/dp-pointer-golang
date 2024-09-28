//write a program that uses the new function to allocate memory for an integer pointer
package main
import "fmt"

func main(){
  p := new(int)
  fmt.Println("Before:",*p)
  *p = 42
  fmt.Println("After:",*p)
}
//The new function allocates memory for an integer and returns a pointer to it. The memory is zero-initialized by default.
