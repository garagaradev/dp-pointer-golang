//demonstrate how to safely handle dereferencing a potentially nil pointer.
package main
import "fmt"

func printValue(ptr *int){
  if ptr == nil {
    fmt.Println("Pointer is nil.")
    return
  }

  fmt.Println("Value:",*ptr)
}

func main(){
  var ptr *int 
  printValue(ptr)

  x := 42
  ptr = &x
  printValue(ptr)
}
