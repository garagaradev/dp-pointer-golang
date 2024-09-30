//demonstrate how to use a pointer to a channel
package main
import "fmt"

func sendData(ch *chan int){
  *ch <- 42
}

func main(){
  ch := make(chan int)
  go sendData(&ch)

  fmt.Println("Received:",<-ch)
}
