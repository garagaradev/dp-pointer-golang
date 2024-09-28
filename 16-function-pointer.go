//write a program that uses a pointer to a function
package main
import "fmt"

func add(a, b int) int {
  return a + b
}

func main(){
  var fptr func(int, int) int
  fptr = add
  fmt.Println(add)
  fmt.Println(fptr)

  result := fptr(3,5)
  fmt.Println(result)
}
