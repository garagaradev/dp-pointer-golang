//use a pointer to a variable shared betweens goroutines.
package main
import (
  "fmt"
  "sync"
) 

func increment(wg *sync.WaitGroup, counter *int){
  defer wg.Done()
  *counter++
}

func main(){
  var wg sync.WaitGroup
  counter := 0

  for i:=0;i<5;i++ {
    wg.Add(1)
    go increment(&wg, &counter)
  }

  wg.Wait()

  fmt.Println("Final Counter:", counter)
}


