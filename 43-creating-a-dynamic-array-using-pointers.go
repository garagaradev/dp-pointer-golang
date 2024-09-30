//demonstrate how to create a dynamic array using pointers
package main
import "fmt"

func createDynamicArray(size int) *[]int{
  arr := make([]int, size)
  return &arr
}

func main(){
  arrPtr := createDynamicArray(5)
  for i:=0;i<5;i++{
    (*arrPtr)[i] = 10 * i
  }

  fmt.Println(*arrPtr)
}
