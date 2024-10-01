//Implement a recursive function to calculate the factorial of a number using pointers.
package main
import "fmt"

func factorial(n *int) int {
  if *n <= 1 {
    return 1
  }
  temp := *n
  *n--
  return temp * factorial(n)
}

func main(){
  num := 5
  fmt.Println("Factorial:", factorial(&num))
}
