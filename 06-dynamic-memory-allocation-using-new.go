//write a program that uses the new function to allocate memory for an integer pointer
package main
import "fmt"

func main(){
  p := new(int)
  fmt.Println("Before:",*p)
  *p = 42
  fmt.Println("After:",*p)
}
