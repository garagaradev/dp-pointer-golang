//show how to iterate over an array using pointer arithmetic
package main
import "fmt"

func printArray(arr *[5]int){
  for i:=0;i<len(arr);i++{
    fmt.Print((*arr)[i:]," ")
  }
  fmt.Println()
}

func main(){
    numbers := [5]int{1,2,3,4,5}
    printArray(&numbers)
  }
