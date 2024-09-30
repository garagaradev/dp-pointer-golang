//write a function that appends an element to a slice using a pointer to the slice
package main
import "fmt"

func appendToSlice(s *[]int, value int){
  *s = append(*s, value)
}

func main(){
  nums := []int{1,2,3}
  appendToSlice(&nums, 4)
  fmt.Println("Slice after append:",nums)
}
