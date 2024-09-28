//write a program that modifies elements of an array using pointers
package main
import "fmt"

func modifyArray(arr *[3]int){
  (*arr)[0] = 10
  (*arr)[1] = 20
  (*arr)[2] = 30
}

func main(){
  arr := [3]int{1,2,3}
  fmt.Println("Before:", arr)
  modifyArray(&arr)
  fmt.Println("After:", arr)
}
