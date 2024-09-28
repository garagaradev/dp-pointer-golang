//write a program that copies a struct containing a pointer
//and show how modifying the pointer affects the original struct
package main
import "fmt"

type Person struct {
  name *string
}

func main(){
  name := "Jack"
  p1 := Person{name: &name}
  p2 := p1

  *p2.name = "Daniel"
  fmt.Println(*p1.name)
  fmt.Println(*p2.name)
}
