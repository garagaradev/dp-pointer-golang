//write a function that adds elements to a slice within a struct using a pointer
package main
import "fmt"

type Collection struct {
  Items []string
}

func addItem(c *Collection, item string){
  c.Items = append(c.Items, item)
}

func main(){
  collection := &Collection{Items:[]string{"Item1"}}
  fmt.Println("Before:", collection.Items)
  addItem(collection, "Item2")
  fmt.Println("After:", collection.Items)
}
