//write a program that sends a pointer to an integer over a channel 
//and modifies the integer at the receiving end.
package main
import "fmt"

func modifyPointer(ch chan *int) {
  p := <-ch
  *p = 50
}

func main(){
  x := 10
  ch := make(chan *int)

  go modifyPointer(ch)
  ch <- &x

  fmt.Println(x)
}
