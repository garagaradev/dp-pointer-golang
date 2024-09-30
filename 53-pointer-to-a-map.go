// demonstrate how to modify a map using a pointer to the map
package main
import "fmt"

func addEntry(m *map[string]int, key string, value int){
  (*m)[key] = value
}

func main(){
  myMap := make(map[string]int)
  addEntry(&myMap, "a", 1)
  addEntry(&myMap, "b", 2)
  fmt.Println("Map entries:", myMap)
}
