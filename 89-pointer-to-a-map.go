//write a function that modifies a map using a pointer
package main
import "fmt"

func addEntry(m *map[string]int, key string, value int){
  (*m)[key] = value
}

func main(){
  myMap := make(map[string]int)
  addEntry(&myMap, "A", 1)
  addEntry(&myMap, "B", 2)

  fmt.Println("Map:", myMap)
}
