//write a function that toggles a boolean value using a pointer.
package main
import "fmt"

func toggle(b *bool){
  *b = !*b
}

func main(){
  flag := true

  fmt.Println("Before toggle:", flag)
  toggle(&flag)
  fmt.Println("After toggle:", flag)
}


