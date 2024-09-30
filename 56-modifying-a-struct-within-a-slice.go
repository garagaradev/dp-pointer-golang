//write a function that modifies a struct in a slice using pointers
package main
import "fmt"

type Item struct {
  name string
  price float64
}

func updateItemPrice(items *[]Item, index int, newPrice float64){
  (*items)[index].price = newPrice
}

func main(){
  items := []Item{{"Apple",1.0},{"Banana",0.5}}
  updateItemPrice(&items, 0, 1.5)
  fmt.Println("Updated items:", items)
}
