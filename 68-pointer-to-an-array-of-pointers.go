//create a program that works with an array of pointers to integers
package main
import "fmt"

func main(){
  nums := [3]*int{new(int), new(int), new(int)}
  *nums[0] = 10
  *nums[1] = 20 
  *nums[2] = 30

  for i, num := range nums {
    fmt.Printf("Value at index %d: %d\n", i, *num)
  }
}
