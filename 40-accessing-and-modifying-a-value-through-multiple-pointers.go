// write a program that modifies a value using multiple levels of pointers.
package main
import "fmt"

func modifyValueThroughPointer(ppp ***int){
  ***ppp = 300
}

func main(){
  x := 100
  p := &x
  pp := &p
  ppp := &pp

  modifyValueThroughPointer(ppp)
  fmt.Println("Modified value:",x)
}
