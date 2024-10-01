//Write a function that takes a pointer to a function, calls it, and returns its result.
package main
import "fmt"

func square(x int) int {
  return x * x
}

func callFunction(fn *func(int) int, value int) int {
  return (*fn)(value)
}

func main(){
  squareFunc := square
  result := callFunction(&squareFunc, 4)
  fmt.Println("Square:", result)
}
