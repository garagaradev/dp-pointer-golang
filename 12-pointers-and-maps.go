//write a program that modifies a value in a map using pointers
package main
import "fmt"

func modifyMap(m map[string]*int){
  if val, exists := m["key"]; exists {
    *val = 100
  }
}

func main(){
  x := 10
  m := map[string]*int{"key":&x}
  fmt.Println("Before:",*m["key"])
  modifyMap(m)
  fmt.Println("After:",*m["key"])
}
