//write a program that swaps two integer values using pointers
package main
import "fmt"

func swap(a,b *int){
  temp := *a
  *a = *b
  *b = temp
}

func main(){
  x, y := 5, 10
  fmt.Println("Before:",x,y)
  swap(&x, &y)
  fmt.Println("After:",x,y)
}
