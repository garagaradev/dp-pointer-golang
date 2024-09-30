// demonstrate how to use pointers to functions
package main
import "fmt"

func add(a,b int)int{
  return a + b
}

func main(){
  funcPtr := add
  result := funcPtr(5,3)
  fmt.Println("Result of addition:", result)
}
