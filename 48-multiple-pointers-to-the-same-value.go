//show how multiple pointers can point to the same value 
//and demonstrate their usage
package main
import "fmt"

func main(){
  x := 50
  p1 := &x
  p2 := &x

  *p1 = 100
  fmt.Println("Value through p1:", *p1)
  fmt.Println("Value through p2:", *p2)

  *p2 = 200
  fmt.Println("Value through p1:",*p1)
  fmt.Println("Value through p2:",*p2)
}
