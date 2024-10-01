//write a function that returns a pointer to another function
package main
import "fmt"

func multiply(a, b int)int{
  return a * b
}

func getMultiplier() func(int, int) int {
  return multiply
}

func main(){
  multiplier:= getMultiplier()
  result := (multiplier)(3,4)

  fmt.Println("Result:",result)
}
