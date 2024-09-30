//write  a function to swap two integers using pointers
package main
import "fmt"

func swap(a *int, b *int){
  *a, *b = *b, *a
}

func main(){
  x := 5
  y := 10
  fmt.Println("Before swap:",x,y)
  swap(&x,&y)
  fmt.Println("After swap:",x,y)
}
