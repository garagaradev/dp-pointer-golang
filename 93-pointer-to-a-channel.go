//create a function that uses a pointer to a channel to send data
package main
import "fmt"

func sendData(ch *chan int, data int){
  *ch <- data
}

func main(){
  ch := make(chan int)
  go sendData(&ch, 42)

  value := <- ch
  fmt.Println("Received:", value)
}
