//create a function that accepts a pointer to another function as a callback
package main
import "fmt"

func add(a, b int) int {
  return a + b
}

func operate(fn *func(int,int)int, x, y int)int{
  return (*fn)(x,y)
}

func main(){
  sum := add
  result := operate(&sum, 3, 4)
  fmt.Println("Sum:", result)
}
