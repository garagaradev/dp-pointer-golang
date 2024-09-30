//use pointer arithmetic to count elements in an array
package main
import "fmt"

func countElements(arr *[5]int) int {
  count := 0
  for i:=0; i<len(arr);i++{
    count++
  }

  return count
}

func main(){
  arr := [5]int{1,2,3,4,5}
  fmt.Println("Number of elements:",countElements(&arr))
}
