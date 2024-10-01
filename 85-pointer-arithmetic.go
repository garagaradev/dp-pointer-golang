//demonstrate a pointer arithmetic using a pointer to an array
package main
import "fmt"

func printArray(arr *[5]int) {
  for i := 0; i < len(arr); i++ {
    fmt.Printf("%d ",(*arr)[i])
  }
  fmt.Println()
}

func main(){
  numbers := [5]int{1,2,3,4,5}
  printArray(&numbers)
}
