//write a program that swaps the values of two integers using pointers
package main
import "fmt"

func swap(a,b *int) {
  *a, *b = *b, *a
}

func main(){
  x, y := 10, 20
  fmt.Println("Before swap:", x, y)
  swap(&x, &y)
  fmt.Println("After swap:", x, y)
}
