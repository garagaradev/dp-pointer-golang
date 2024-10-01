//create a function that modifies an array using a pointer
package main
import "fmt"

func modifyArray(arr *[3]int){
  arr[0] = 100
}

func main(){
  numbers := [3]int{1,2,3}
  fmt.Println("Before:", numbers)
  modifyArray(&numbers)
  fmt.Println("After:", numbers)
}
