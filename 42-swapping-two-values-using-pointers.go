//implement a function that swaps two integers using pointers
package main
import "fmt"

func swap(a *int, b *int){
  temp := *a
  *a = *b
  *b = temp
}

func main(){
  x := 10
  y := 20
  swap(&x, &y)
  fmt.Println("swapped values: x =",x,", y =",y)
}
