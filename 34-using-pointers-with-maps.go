//demonstrate how to store pointers inside a map and 
//manipulate the underlying values
package main
import "fmt"

func main(){
  myMap := make(map[string]*int)

  a := 10
  b := 20

  myMap["first"] = &a
  myMap["second"] = &b

  *myMap["first"] = 100
  *myMap["second"] = 200

  fmt.Println("first:",*myMap["first"])
  fmt.Println("second:",*myMap["second"])
}
