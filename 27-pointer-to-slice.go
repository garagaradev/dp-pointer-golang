// write a program to modify a slice through a pointer to the slice
package main
import "fmt"

func modifySlice(s *[]int){
  *s = append(*s, 4,5)
}

func main(){
  slice := []int{1,2,3}
  modifySlice(&slice)
  fmt.Println(slice)
}
