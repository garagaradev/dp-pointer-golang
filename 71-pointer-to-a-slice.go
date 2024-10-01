//write a function that modifies a slice using a pointer
package main
import "fmt"

func modifySlice(s *[]int){
  (*s)[0] = 100
}

func main(){
  numbers := []int{1,2,3}
  fmt.Println("Before:",numbers)
  modifySlice(&numbers)
  fmt.Println("After:", numbers)
}
