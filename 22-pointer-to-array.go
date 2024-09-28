//write a program that uses a pointer to an array and modifies its elements
package main
import "fmt"

func modifyArray(arr *[3]int){
  for i := range arr {
    arr[i] *= 2
  }
}

func main(){
  arr := [3]int{1,2,3}
  fmt.Println("initial array:",arr)
  modifyArray(&arr)
  fmt.Println("current array:",arr)
}
