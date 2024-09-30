//Demonstrate the usage of pointers with slices
package main
import "fmt"

func modifySlice(s *[]int){
  (*s)[0] = 100
}

func main(){
  nums := []int{1,2,3}
  fmt.Println("original slice:",nums)
  modifySlice(&nums)
  fmt.Println("modified slice:",nums)
}
