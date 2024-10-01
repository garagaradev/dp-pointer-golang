//pointer to a struct with embedded types
package main
import "fmt"

type Base struct {
  Value int
}

type Derived struct {
  Base
}

func modifyBase(b *Base, newValue int){
  b.Value = newValue
}

func main(){
  derived := &Derived{Base{Value:10}}
  fmt.Println("Before:", derived.Value)
  modifyBase(&derived.Base, 100)
  fmt.Println("After:", derived.Value)
}
