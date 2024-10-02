//write a recursive function that uses a pointer to calculate factorial.
package main
import "fmt"

func factorial(n int, result *int){
  if n == 0 {
    *result = 1
    return
  }
  factorial(n-1, result)
  *result *= n 
}

func main(){
  var result int
  factorial(5, &result)
  fmt.Println("Factorial:", result)
}
