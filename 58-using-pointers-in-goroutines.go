//create a goroutine that modifies a value using a pointer
package main
import "fmt"

func updateValue(val *int){
  *val = 100
}

func main(){
  num := 0
  go updateValue(&num)

  //adding a small delay to allow goroutine to finish
  //in production code,use sync.WaitGroup or other methods
  fmt.Scanln()

  fmt.Println("Updated value:", num)
}
