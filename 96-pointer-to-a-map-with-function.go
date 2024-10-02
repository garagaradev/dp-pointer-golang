//create a function that takes a pointer to a map and performs operations on it.
package main
import "fmt"

func modifyMap(m *map[string]int){
  (*m)["A"] = 10
  (*m)["B"] = 20
}

func main(){
  myMap := make(map[string]int)
  modifyMap(&myMap)
  fmt.Println("Modified map:", myMap)
}
