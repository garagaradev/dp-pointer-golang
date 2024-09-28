//write a program that attempts to add an integer to a pointer and observe the error
package main
import "fmt"

func main(){
  x := 5
  p := &x

  p++
  fmt.Println(p)
}
