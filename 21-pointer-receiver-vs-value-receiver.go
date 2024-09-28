// demonstrate the difference between pointer receiver and value receiver in methods
package main
import "fmt"

type Counter struct {
  count int
}

func (c *Counter) IncrementPointer(){
  c.count++
}

func (c Counter) IncrementValue(){
  c.count++
}

func main(){
  c := Counter{count:10}

  c.IncrementValue()
  fmt.Println("After IncrementValue:",c.count)

  c.IncrementPointer()
  fmt.Println("After IncrementPointer:",c.count)


}
