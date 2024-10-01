//create multiple goroutines that send data to a shared channel using a pointer to the channel
package main
import (
  "fmt"
  "sync"
)

func sendData(ch *chan int, wg *sync.WaitGroup, value int){
  defer wg.Done()
  *ch <- value
}

func main(){
  ch := make(chan int)
  var wg sync.WaitGroup

  wg.Add(3)
  go sendData(&ch, &wg, 1)
  go sendData(&ch, &wg, 2)
  go sendData(&ch, &wg, 3)

  go func(){
    wg.Wait()
    close(ch)
  }()

  for value := range ch {
    fmt.Println("Received:",value)
  }
}
